package main

func (d *Daemon) setRoutes() {
	d.Engine.GET("/groups", d.getGroups)

	allGroupRouter := d.Engine.Group("/groups/")
	{
		allGroupRouter.GET("", d.getGroups)

		oneGroupRouter := allGroupRouter.Group("/:group")
		{
			oneGroupRouter.GET("", d.getGroup)
			oneGroupRouter.POST("", d.postGroup)
			oneGroupRouter.DELETE("", d.deleteGroup)

			oneGroupConnectorRouter := oneGroupRouter.Group("/connectors")
			{
				oneGroupConnectorRouter.GET("", d.getConnectors)
				oneGroupConnectorRouter.GET("/:connector", d.getConnector)
				oneGroupConnectorRouter.POST("/:connector", d.createConnector)
			}
		}
	}

}
