package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samalba/dockerclient"
	"gopkg.in/redis.v3"
	"strings"
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
    d.Engine.GET("/groups/:group/connectors", d.getConnectors)
}

func (d *Daemon) getGroups(c *gin.Context) {
    groups := d.Intools.getGroups()
    c.String(200, "-> %s", strings.Join(groups, ";"))
}

func (d *Daemon) getConnectors(c *gin.Context) {
    group := c.Param("group")
    connectors := d.Intools.getConnectors(group)
    c.String(200, "-> %s", strings.Join(connectors, ";"))
}



