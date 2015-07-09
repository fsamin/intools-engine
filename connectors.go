package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (d *Daemon) getConnectors(c *gin.Context) {
	group := c.Param("group")
	connectors := d.Intools.getConnectors(group)
	c.String(200, "-> %s", strings.Join(connectors, ";"))
}
