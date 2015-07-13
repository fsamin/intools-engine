package main

import (
	"github.com/gin-gonic/gin"
)

func (d *Daemon) getGroups(c *gin.Context) {
	groups := Intools.GetGroups(false)
	c.JSON(200, groups)
}

func (d *Daemon) getGroup(c *gin.Context) {
	group := c.Param("group")
	g := Intools.GetGroup(group, false)
	if g == nil {
		c.String(404, "")
	} else {
		c.JSON(200, g)
	}
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
