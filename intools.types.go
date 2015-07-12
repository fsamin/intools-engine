package main

import (
	"encoding/json"
	"github.com/robfig/cron"
	"github.com/samalba/dockerclient"
	"gopkg.in/redis.v3"
	"time"
)

type IntoolsEngine struct {
	DockerClient dockerclient.DockerClient
	DockerHost   string
	RedisClient  redis.Client
	Cron         *cron.Cron
}

type Connector struct {
	Engine          *IntoolsEngine
	Group           string
	Name            string
	ContainerConfig *dockerclient.ContainerConfig
	Timeout         int
	Refresh         int
}

func (c *Connector) GetContainerName() string {
	return c.Group + "-" + c.Name
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
