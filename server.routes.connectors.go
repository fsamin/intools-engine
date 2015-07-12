package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (d *Daemon) getConnectors(c *gin.Context) {
	group := c.Param("group")
	connectors := d.Intools.GetConnectors(group)
	c.String(200, "-> %s", strings.Join(connectors, ";"))
}

func (d *Daemon) getConnector(c *gin.Context) {
    group := c.Param("group")
    connector := c.Param("connector")

    Debug.Printf("Searching for %s:%s", group, connector)

    conn, err := d.Intools.GetConnector(group, connector)
    if err != nil {
        c.String(404, err.Error())
    } else {
        c.String(200, conn.GetJSON())
    }
}

func (d *Daemon) createConnector(c *gin.Context) {
    group := c.Param("group")
    connector := c.Param("connector")

    var conn Connector
    c.Bind(&conn)
    conn.Group = group
    conn.Name = connector

    d.Intools.SaveConnector(&conn)

    conn.Engine = d.Intools
    conn.InitSchedule()

    c.String(200, conn.GetJSON())
}


