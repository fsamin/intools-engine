package main

import (
	"github.com/gin-gonic/gin"
)

func (d *Daemon) getConnectors(c *gin.Context) {
	group := c.Param("group")
	connectors := Intools.GetConnectors(group)
	c.JSON(200, connectors)
}

func (d *Daemon) getConnector(c *gin.Context) {
	group := c.Param("group")
	connector := c.Param("connector")

	Debug.Printf("Searching for %s:%s", group, connector)

	conn, err := Intools.GetConnector(group, connector)
	if err != nil {
		c.String(404, err.Error())
	} else {
		c.JSON(200, conn)
	}
}

func (d *Daemon) createConnector(c *gin.Context) {
	group := c.Param("group")
	connector := c.Param("connector")

	var conn Connector
	c.Bind(&conn)
	conn.Group = group
	conn.Name = connector

	Intools.SaveConnector(&conn)
	Intools.InitSchedule(&conn)

	c.JSON(200, conn)
}
