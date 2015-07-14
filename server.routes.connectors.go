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

func (d *Daemon) execConnector(c *gin.Context) {
	group := c.Param("group")
	connector := c.Param("connector")

	Debug.Printf("Searching for %s:%s", group, connector)

	conn, err := Intools.GetConnector(group, connector)
	if err != nil {
		c.String(404, err.Error())
	} else {
		executor, err := Intools.Exec(conn)
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, executor)
		}
	}
}

func (d *Daemon) getConnectorExecutor(c *gin.Context) {
	group := c.Param("group")
	connector := c.Param("connector")
	conn, err := Intools.GetConnector(group, connector)
	if err != nil {
		c.String(404, err.Error())
	} else {
		exec := Intools.GetLastConnectorExecutor(conn)
		if exec == nil {
			c.String(404, "no executor found")
		} else {
			c.JSON(200, exec)
		}
	}
}

func (d *Daemon) getConnectorResult(c *gin.Context) {
	group := c.Param("group")
	connector := c.Param("connector")
	conn, err := Intools.GetConnector(group, connector)
	if err != nil {
		c.String(404, err.Error())
	} else {
		exec := Intools.GetLastConnectorExecutor(conn)
		if exec == nil {
			c.String(404, "no result found")
		} else {
			if exec.Valid {
				c.JSON(200, exec.JsonStdout)
			} else {
				c.String(404, "invalid result")
			}
		}
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
