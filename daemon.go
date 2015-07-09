package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samalba/dockerclient"
	"gopkg.in/redis.v3"
)

type Daemon struct {
	Port      int
	Engine    *gin.Engine
	DebugMode bool
	Intools   *IntoolsEngine
}

func newDaemon(port int, debug bool, dockerClient *dockerclient.DockerClient, dockerHost string, redisClient *redis.Client) *Daemon {
	if debug {
		Debug.Println("Initializing daemon in debug mode")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	intools := NewIntoolsEngine(dockerClient, dockerHost, *redisClient)
	daemon := Daemon{port, engine, debug, intools}
	return &daemon
}

func (d *Daemon) run() {
	d.Engine.Run(fmt.Sprintf("0.0.0.0:%d", d.Port))
}

func (d *Daemon) setRoutes() {
	d.Engine.GET("/groups", d.getGroups)

	allGroupRouter := d.Engine.Group("/groups/")
	{
		allGroupRouter.GET("", d.getGroups)

		onGroupRouteur := allGroupRouter.Group("/:group")
		{
			onGroupRouteur.GET("", d.getGroup)
			onGroupRouteur.POST("", d.postGroup)
			onGroupRouteur.DELETE("", d.deleteGroup)
			onGroupRouteur.GET("/connectors", d.getConnectors)
		}
	}

}
