package intools

import (
	"github.com/robfig/cron"
	"github.com/samalba/dockerclient"
	"gopkg.in/redis.v3"
)

type IntoolsEngine struct {
	DockerClient dockerclient.Client
	DockerHost   string
	RedisClient  redis.Client
	Cron         *cron.Cron
}

var (
	Engine *IntoolsEngine
)
