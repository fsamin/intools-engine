package main

import (
	"encoding/json"
	"github.com/samalba/dockerclient"
	"gopkg.in/redis.v3"
	"time"
)

type IntoolsEngine struct {
	DockerClient dockerclient.Client
	DockerHost   string
	RedisClient  redis.Client
}

type Connector struct {
	Group           string
	Name            string
	ContainerConfig *dockerclient.ContainerConfig
	Timeout         int
}

func (c *Connector) Init(image string, timeout int, cmd []string) {
    c.ContainerConfig = &dockerclient.ContainerConfig{
        Image:        image,
        Cmd:          cmd,
        AttachStdin:  false,
        AttachStdout: false,
        AttachStderr: false,
        Tty:          false,
    }
    if timeout != 0 {
        c.Timeout = timeout
    }
}

func (c *Connector) GetContainerName() string {
	return c.Group + "-" + c.Name
}

func (c *Connector) GetKey() string {
	return "intools:groups:" + c.Group + ":connectors:" + c.Name
}

func (c *Connector) GetExecutorKey(e *Executor) string {
	return "intools:groups:" + c.Group + ":connectors:" + c.Name + ":executors"
}

func (c *Connector) GetResultKey(e *Executor) string {
	return "intools:groups:" + c.Group + ":connectors:" + c.Name + ":results"
}

func (c *Connector) GetJSON() string {
	b, err := json.Marshal(c)
	if err != nil {
		Error.Println(err)
		return ""
	}
	return string(b[:])
}



type Executor struct {
	ContainerId string
	Host        string
	Running     bool
	Terminated  bool
	ExitCode    int
	Stdout      string
	JsonStdout  *map[string]interface{}
	Stderr      string
	StartedAt   time.Time
	FinishedAt  time.Time
	Valid       bool
}

func (e *Executor) GetJSON() string {
	b, err := json.Marshal(e)
	if err != nil {
		Error.Println(err)
		return ""
	}
	return string(b[:])
}

func (e *Executor) GetResult() string {
	b, err := json.Marshal(e.JsonStdout)
	if err != nil {
		Error.Println(err)
		return ""
	}
	return string(b[:])
}
