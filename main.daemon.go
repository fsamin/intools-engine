package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/samalba/dockerclient"
	"gopkg.in/redis.v3"
)

type Daemon struct {
	Port      int
	Engine    *gin.Engine
	DebugMode bool
}

func newDaemon(port int, debug bool, dockerClient *dockerclient.DockerClient, dockerHost string, redisClient *redis.Client) *Daemon {
	if debug {
		Debug.Println("Initializing daemon in debug mode")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	cron := cron.New()
	Intools = &IntoolsEngine{*dockerClient, dockerHost, *redisClient, cron}
	daemon := &Daemon{port, engine, debug}
	return daemon
}

func (d *Daemon) run() {
	go func() {
		Intools.Reload()
		Intools.Cron.Start()
	}()
	d.Engine.Run(fmt.Sprintf("0.0.0.0:%d", d.Port))
}
