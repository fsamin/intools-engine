package groups

import (
	"github.com/gin-gonic/gin"
)

func ControllerGetGroups(c *gin.Context) {
	groups := GetGroups(false)
	c.JSON(200, groups)
}

func ControllerGetGroup(c *gin.Context) {
	group := c.Param("group")
	g := GetGroup(group, false)
	if g == nil {
		c.String(404, "")
	} else {
		c.JSON(200, g)
	}
}

func ControllerPostGroup(c *gin.Context) {
	group := c.Param("group")
	created, err := CreateGroup(group)
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

func ControllerDeleteGroup(c *gin.Context) {
	group := c.Param("group")
	err := DeleteGroup(group)
	if err != nil {
		c.String(500, err.Error())
	} else {
		c.String(200, "%s deleted", group)
	}
}
