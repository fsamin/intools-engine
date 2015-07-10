package main

func (d *Daemon) setRoutes() {
	d.Engine.GET("/groups", d.getGroups)

	allGroupRouter := d.Engine.Group("/groups/")
	{
		allGroupRouter.GET("", d.getGroups)

		onGroupRouteur := allGroupRouter.Group("/:group")
		{
			onGroupRouteur.GET("", d.getGroup)
			onGroupRouteur.POST("", d.postGroup)
			onGroupRouteur.DELETE("", d.deleteGroup)
			onGroupRouteur.GET("/connectors", d.getConnectors)
		}
	}

}
