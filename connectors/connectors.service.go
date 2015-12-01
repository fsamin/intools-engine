package connectors

import (
	"encoding/json"
	"fmt"
	"github.com/fsamin/intools-engine/common/logs"
	"github.com/fsamin/intools-engine/common/utils"
	"github.com/fsamin/intools-engine/common/websocket"
	"github.com/fsamin/intools-engine/executors"
	"github.com/fsamin/intools-engine/intools"
	"github.com/samalba/dockerclient"
	"sync"
	"time"
)

func InitSchedule(c *Connector) {
	if intools.Engine.GetCron() != nil {
		crontab := fmt.Sprintf("@every %ds", c.Refresh)
		logs.Debug.Printf("Schedule %s:%s %s", c.Group, c.Name, crontab)
		intools.Engine.GetCron().AddJob(crontab, c)
	}
}

func Exec(connector *Connector) (*executors.Executor, error) {
	executor := &executors.Executor{}

	//Saving connector to redis
	go SaveConnector(connector)

	//Get all containers
	containers, err := intools.Engine.GetDockerClient().ListContainers(true, false, "")
	if err != nil {
		logs.Error.Println(err)
		return nil, err
	}

	//Searching for the container with the same name
	containerExists := false
	previousContainerId := "-1"
	for _, c := range containers {
		for _, n := range c.Names {
			if n == fmt.Sprintf("/%s", connector.GetContainerName()) {
				containerExists = true
				previousContainerId = c.Id
			}
		}
	}

	//If it exists, remove it
	if containerExists {
		logs.Trace.Printf("Removing container %s [/%s]", previousContainerId[:11], connector.GetContainerName())
		err := intools.Engine.GetDockerClient().RemoveContainer(previousContainerId, true, true)
		if err != nil {
			logs.Error.Println("Cannot remove container " + previousContainerId[:11])
			logs.Error.Println(err)
			return nil, err
		}
	}

	//Create container
	ContainerId, err := intools.Engine.GetDockerClient().CreateContainer(connector.ContainerConfig, connector.GetContainerName())
	if err != nil {
		logs.Error.Println("Cannot create container " + connector.GetContainerName())
		logs.Error.Println(err)
		return nil, err
	}
	//Save the short ContainerId
	executor.ContainerId = ContainerId[:11]
	executor.Host = intools.Engine.GetDockerHost()

	logs.Trace.Printf("%s [/%s] successfully created", executor.ContainerId, connector.GetContainerName())
	hostConfig := &dockerclient.HostConfig{}

	//Prepare the waiting group to sync execution of the container
	var wg sync.WaitGroup
	wg.Add(1)

	//Start the container
	err = intools.Engine.GetDockerClient().StartContainer(ContainerId, hostConfig)
	if err != nil {
		logs.Error.Println("Cannot start container " + executor.ContainerId)
		logs.Error.Println(err)
		return nil, err
	}

	logs.Trace.Printf("%s [/%s] successfully started", executor.ContainerId, connector.GetContainerName())
	logs.Debug.Println(executor.ContainerId + " will be stopped after " + fmt.Sprint(connector.Timeout) + " seconds")
	//Trigger stop of the container after the timeout
	intools.Engine.GetDockerClient().StopContainer(ContainerId, connector.Timeout)

	//Wait for the end of the execution of the container
	for {
		//Each time inspect the container
		inspect, err := intools.Engine.GetDockerClient().InspectContainer(ContainerId)
		if err != nil {
			logs.Error.Println("Cannot inpect container " + executor.ContainerId)
			logs.Error.Println(err)
			return executor, err
		}
		if !inspect.State.Running {
			//When the container is not running
			logs.Debug.Println(executor.ContainerId + " is stopped")
			executor.Running = false
			executor.Terminated = true
			executor.ExitCode = inspect.State.ExitCode
			executor.StartedAt = inspect.State.StartedAt
			executor.FinishedAt = inspect.State.FinishedAt
			//Trigger next part of the waiting group
			wg.Done()
			//Exit from the waiting loop
			break
		} else {
			//Wait
			logs.Debug.Println(executor.ContainerId + " is running...")
			time.Sleep(5 * time.Second)
		}
	}

	//Next part : after the container has been executed
	wg.Wait()

	logStdOutOptions := &dockerclient.LogOptions{
		Follow:     true,
		Stdout:     true,
		Stderr:     false,
		Timestamps: false,
		Tail:       0,
	}

	logStdErrOptions := &dockerclient.LogOptions{
		Follow:     true,
		Stdout:     false,
		Stderr:     true,
		Timestamps: false,
		Tail:       0,
	}

	//Get the stdout and stderr
	logsStdOutReader, err := intools.Engine.GetDockerClient().ContainerLogs(ContainerId, logStdOutOptions)
	logsStdErrReader, err := intools.Engine.GetDockerClient().ContainerLogs(ContainerId, logStdErrOptions)

	if err != nil {
		logs.Error.Println("-cannot read logs from server")
	} else {
		containerLogs, err := utils.ReadLogs(logsStdOutReader)
		if err != nil {
			return executor, err
		} else {
			executor.Stdout = containerLogs
			executor.JsonStdout = new(map[string]interface{})
			errJsonStdOut := json.Unmarshal([]byte(executor.Stdout), executor.JsonStdout)
			executor.Valid = true
			if errJsonStdOut != nil {
				logs.Warning.Printf("Unable to parse stdout from container %s", executor.ContainerId)
				logs.Warning.Println(errJsonStdOut)
			}
		}
		containerLogs, err = utils.ReadLogs(logsStdErrReader)
		if err != nil {
			return executor, err
		} else {
			executor.Stderr = containerLogs
		}
	}

	// Broadcast result to registered clients
	websocket.Notify(connector.Group, connector.Name, executor.JsonStdout)

	//Save result to redis
	defer SaveExecutor(connector, executor)

	return executor, nil
}
