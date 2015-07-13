package main

import (
	"encoding/json"
	"fmt"
	"github.com/samalba/dockerclient"
	"sync"
	"time"
)

func (eng *IntoolsEngine) NewConnector(group string, name string) *Connector {
	conn := &Connector{group, name, nil, 15, 60}
	return conn
}

func (eng *IntoolsEngine) Reload() {
	groups := eng.GetGroups(true)
	for _, group := range groups {
		Trace.Printf("%s - Reloading group", group.Name)
		for _, connector := range group.Connectors {
			Trace.Printf("%s:%s - Reloading connector", group.Name, connector.Name)
			eng.InitSchedule(&connector)
		}
	}
}

func (e *IntoolsEngine) InitSchedule(c *Connector) {
	if e.Cron != nil {
		crontab := fmt.Sprintf("@every %dm", c.Refresh)
		Debug.Printf("Schedule %s:%s %s", c.Group, c.Name, crontab)
		e.Cron.AddJob(crontab, c)
	}
}

func (c *Connector) Init(image string, timeout int, refresh int, cmd []string) {
	if c.ContainerConfig == nil {
		c.ContainerConfig = &dockerclient.ContainerConfig{
			Image:        image,
			Cmd:          cmd,
			AttachStdin:  false,
			AttachStdout: false,
			AttachStderr: false,
			Tty:          false,
		}
	}

	if timeout != 0 {
		c.Timeout = timeout
	}
	if refresh != 0 {
		c.Refresh = refresh
	}
}

func (c *Connector) Run() {
	//TODO : Should not run error, or invalid connector ?
	Debug.Printf("Run Connector %s:%s", c.Group, c.Name)
	Intools.Exec(c)
}

func (e *IntoolsEngine) Exec(connector *Connector) (*Executor, error) {
	executor := &Executor{}

	//Saving connector to redis
	go e.SaveConnector(connector)

	//Get all containers
	containers, err := e.DockerClient.ListContainers(true, false, "")
	if err != nil {
		Error.Println(err)
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
		Trace.Printf("Removing container %s [/%s]", previousContainerId[:11], connector.GetContainerName())
		err := e.DockerClient.RemoveContainer(previousContainerId, true, true)
		if err != nil {
			Error.Println("Cannot remove container " + previousContainerId[:11])
			Error.Println(err)
			return nil, err
		}
	}

	//Create container
	ContainerId, err := e.DockerClient.CreateContainer(connector.ContainerConfig, connector.GetContainerName())
	if err != nil {
		Error.Println("Cannot create container " + connector.GetContainerName())
		Error.Println(err)
		return nil, err
	}
	//Save the short ContainerId
	executor.ContainerId = ContainerId[:11]
	executor.Host = e.DockerHost

	Trace.Printf("%s [/%s] successfully created", executor.ContainerId, connector.GetContainerName())
	hostConfig := &dockerclient.HostConfig{}

	//Prepare the waiting group to sync execution of the container
	var wg sync.WaitGroup
	wg.Add(1)

	//Start the container
	err = e.DockerClient.StartContainer(ContainerId, hostConfig)
	if err != nil {
		Error.Println("Cannot start container " + executor.ContainerId)
		Error.Println(err)
		return nil, err
	}

	Trace.Printf("%s [/%s] successfully started", executor.ContainerId, connector.GetContainerName())
	Debug.Println(executor.ContainerId + " will be stopped after " + fmt.Sprint(connector.Timeout) + " seconds")
	//Trigger stop of the container after the timeout
	e.DockerClient.StopContainer(ContainerId, connector.Timeout)

	//Wait for the end of the execution of the container
	for {
		//Each time inspect the container
		inspect, err := e.DockerClient.InspectContainer(ContainerId)
		if err != nil {
			Error.Println("Cannot inpect container " + executor.ContainerId)
			Error.Println(err)
			return executor, err
		}
		if !inspect.State.Running {
			//When the container is not running
			Debug.Println(executor.ContainerId + " is stopped")
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
			Debug.Println(executor.ContainerId + " is running...")
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
	logsStdOutReader, err := e.DockerClient.ContainerLogs(ContainerId, logStdOutOptions)
	logsStdErrReader, err := e.DockerClient.ContainerLogs(ContainerId, logStdErrOptions)

	if err != nil {
		Error.Println("-cannot read logs from server")
	} else {
		logs, err := ReadLogs(logsStdOutReader)
		if err != nil {
			return executor, err
		} else {
			executor.Stdout = logs
			executor.JsonStdout = new(map[string]interface{})
			errJsonStdOut := json.Unmarshal([]byte(executor.Stdout), executor.JsonStdout)
			executor.Valid = true
			if errJsonStdOut != nil {
				Warning.Printf("Unable to parse stdout from container %s", executor.ContainerId)
				Warning.Println(errJsonStdOut)
			}
		}
		logs, err = ReadLogs(logsStdErrReader)
		if err != nil {
			return executor, err
		} else {
			executor.Stderr = logs
		}
	}

	//Save result to redis
	defer e.SaveExecutor(connector, executor)

	return executor, nil
}
