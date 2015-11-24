package tests

import (
	"github.com/samalba/dockerclient"
	"github.com/fsamin/intools-engine/intools"
)

type IntoolsEngineMock struct {
	DockerClient dockerclient.Client
	DockerHost   string
	RedisClient  intools.RedisWrapper
	Cron         intools.CronWrapper
}

func (e *IntoolsEngineMock) GetDockerClient() dockerclient.Client {
	return e.DockerClient
}

func (e *IntoolsEngineMock) GetDockerHost() string {
	return e.DockerHost
}

func (e *IntoolsEngineMock) GetRedisClient() intools.RedisWrapper {
	return e.RedisClient
}

func (e *IntoolsEngineMock) GetCron() intools.CronWrapper {
	return e.GetCron()
}
