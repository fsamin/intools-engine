package server

import (
	"fmt"
	"github.com/fsamin/intools-engine/common/logs"
	"github.com/fsamin/intools-engine/common/websocket"
	"github.com/fsamin/intools-engine/connectors"
	"github.com/fsamin/intools-engine/groups"
	"github.com/fsamin/intools-engine/intools"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/samalba/dockerclient"
	"gopkg.in/redis.v3"

	"github.com/gin-gonic/contrib/expvar"
)

type Daemon struct {
	Port      int
	Engine    *gin.Engine
	DebugMode bool
}

func NewDaemon(port int, debug bool, dockerClient *dockerclient.DockerClient, dockerHost string, redisClient *redis.Client) *Daemon {

	engine := gin.Default()
	if debug {
		logs.Debug.Println("Initializing daemon in debug mode")
		gin.SetMode(gin.DebugMode)
		engine.LoadHTMLFiles("index.html")
		engine.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", nil)
		})
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	cron := cron.New()
	intools.Engine = &intools.IntoolsEngineImpl{dockerClient, dockerHost, redisClient, cron}
	daemon := &Daemon{port, engine, debug}

	length := groups.GetGroupsLength()
	websocket.InitChannel(length)
	return daemon
}

func (d *Daemon) Run() {
	go func() {
		groups.Reload()
		intools.Engine.GetCron().Start()
	}()
	d.Engine.Run(fmt.Sprintf("0.0.0.0:%d", d.Port))
}

func (d *Daemon) SetRoutes() {
	d.Engine.GET("/websocket", websocket.GetWS)
	d.Engine.GET("/debug/vars", expvar.Handler())
	d.Engine.GET("/groups", groups.ControllerGetGroups)

	allGroupRouter := d.Engine.Group("/groups/")
	{
		allGroupRouter.GET("", groups.ControllerGetGroups)

		oneGroupRouter := allGroupRouter.Group("/:group")
		{
			oneGroupRouter.GET("", groups.ControllerGetGroup)
			oneGroupRouter.POST("", groups.ControllerPostGroup)
			oneGroupRouter.DELETE("", groups.ControllerDeleteGroup)

			oneGroupConnectorRouter := oneGroupRouter.Group("/connectors")
			{
				oneGroupConnectorRouter.GET("", connectors.ControllerGetConnectors)
				oneGroupConnectorRouter.GET("/:connector", connectors.ControllerGetConnector)
				oneGroupConnectorRouter.POST("/:connector", connectors.ControllerCreateConnector)
				oneGroupConnectorRouter.GET("/:connector/refresh", connectors.ControllerExecConnector)
				oneGroupConnectorRouter.GET("/:connector/result", connectors.ControllerGetConnectorResult)
				oneGroupConnectorRouter.GET("/:connector/exec", connectors.ControllerGetConnectorExecutor)
			}
		}
	}

}
