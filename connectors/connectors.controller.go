package connectors

import (
	"github.com/fsamin/intools-engine/common/logs"
	"github.com/gin-gonic/gin"
)

func ControllerGetConnectors(c *gin.Context) {
	group := c.Param("group")
	connectors := GetConnectors(group)
	c.JSON(200, connectors)
}

func ControllerGetConnector(c *gin.Context) {
	group := c.Param("group")
	connector := c.Param("connector")

	logs.Debug.Printf("Searching for %s:%s", group, connector)

	conn, err := GetConnector(group, connector)
	if err != nil {
		c.String(404, err.Error())
	} else {
		c.JSON(200, conn)
	}
}

func ControllerExecConnector(c *gin.Context) {
	group := c.Param("group")
	connector := c.Param("connector")

	logs.Debug.Printf("Searching for %s:%s", group, connector)

	conn, err := GetConnector(group, connector)
	if err != nil {
		c.String(404, err.Error())
	} else {
		executor, err := Exec(conn)
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.JSON(200, executor)
		}
	}
}

func ControllerGetConnectorExecutor(c *gin.Context) {
	group := c.Param("group")
	connector := c.Param("connector")
	conn, err := GetConnector(group, connector)
	if err != nil {
		c.String(404, err.Error())
	} else {
		exec := GetLastConnectorExecutor(conn)
		if exec == nil {
			c.String(404, "no executor found")
		} else {
			c.JSON(200, exec)
		}
	}
}

func ControllerGetConnectorResult(c *gin.Context) {
	group := c.Param("group")
	connector := c.Param("connector")
	conn, err := GetConnector(group, connector)
	if err != nil {
		c.String(404, err.Error())
	} else {
		exec := GetLastConnectorExecutor(conn)
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

func ControllerCreateConnector(c *gin.Context) {
	group := c.Param("group")
	connector := c.Param("connector")

	var conn Connector
	c.Bind(&conn)
	conn.Group = group
	conn.Name = connector

	SaveConnector(&conn)
	InitSchedule(&conn)

	c.JSON(200, conn)
}
