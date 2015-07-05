package main

func (e *IntoolsEngine) saveConnector(c *Connector) {
	Debug.Printf("Saving %s to redis", c.GetContainerName())
	_ = e.RedisClient.Set(c.GetKey(), c.GetJSON(), 0)

}

func (e *IntoolsEngine) saveExecutor(c *Connector, exec *Executor) {
	Debug.Printf("Saving %s:%s to redis", c.GetContainerName(), exec.ContainerId)
	_ = e.RedisClient.LPush(c.GetExecutorKey(exec), exec.GetJSON())
	if exec.Valid {
		_ = e.RedisClient.LPush(c.GetResultKey(exec), exec.GetResult())
	}
}
