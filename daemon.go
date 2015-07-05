package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Daemon struct {
	Port      int
	Engine    *gin.Engine
	DebugMode bool
}

func newDaemon(port int, debug bool) *Daemon {
	if debug {
		Debug.Println("Initializing daemon in debug mode")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	daemon := Daemon{port, engine, debug}
	return &daemon
}

func (d *Daemon) setRoutes() {
	d.Engine.GET("/connectors/:group", func(c *gin.Context) {
		group := c.Param("group")
		c.String(200, "Hello %s", group)
	})

}

func (d *Daemon) run() {
	d.Engine.Run(fmt.Sprintf("0.0.0.0:%d", d.Port))
}
