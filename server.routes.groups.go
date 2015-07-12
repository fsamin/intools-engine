package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (d *Daemon) getGroups(c *gin.Context) {
	groups := Intools.GetGroups()
	c.String(200, "-> %s", strings.Join(groups, ";"))
}

func (d *Daemon) getGroup(c *gin.Context) {
	groups := Intools.GetGroups()
	c.String(200, "-> %s", strings.Join(groups, ";"))
}

func (d *Daemon) postGroup(c *gin.Context) {
	group := c.Param("group")
	created, err := Intools.CreateGroup(group)
	if err != nil {
		c.String(500, err.Error())
	} else {
		if created {
			c.String(201, "%s created", group)
		} else {
			c.String(200, "%s already exists", group)
		}
	}
}

func (d *Daemon) deleteGroup(c *gin.Context) {
	group := c.Param("group")
	err := Intools.DeleteGroup(group)
	if err != nil {
		c.String(500, err.Error())
	} else {
		c.String(200, "%s deleted", group)
	}
}
