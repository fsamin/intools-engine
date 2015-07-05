package main
import "fmt"

func (c *Connector) GetRedisGroupKey() string {
    return "intools:groups"
}

func (c *Connector) GetRedisConnectorsKey() string {
    return "intools:groups:" + c.Group + ":connectors"
}

func (c *Connector) GetRedisConnectorKey() string {
    return "intools:groups:" + c.Group + ":connectors:" + c.Name
}

func (c *Connector) GetRedisExecutorKey(e *Executor) string {
    return "intools:groups:" + c.Group + ":connectors:" + c.Name + ":executors"
}

func (c *Connector) GetRedisResultKey(e *Executor) string {
    return "intools:groups:" + c.Group + ":connectors:" + c.Name + ":results"
}

func (e *IntoolsEngine) saveConnector(c *Connector) {
    Debug.Printf("Saving %s to redis", c.Group)
	_ = e.RedisClient.LPush(c.GetRedisGroupKey(), c.Group)
    _ = e.RedisClient.LPush(c.GetRedisConnectorsKey(), c.Name)
	Debug.Printf("Saving %s to redis", c.GetContainerName())
	_ = e.RedisClient.Set(c.GetRedisConnectorKey(), c.GetJSON(), 0)
    return
}

func (e *IntoolsEngine) saveExecutor(c *Connector, exec *Executor) {
	Debug.Printf("Saving %s:%s to redis", c.GetContainerName(), exec.ContainerId)
	_ = e.RedisClient.LPush(c.GetRedisExecutorKey(exec), exec.GetJSON())
	if exec.Valid {
		_ = e.RedisClient.LPush(c.GetRedisResultKey(exec), exec.GetResult())
	}
    return
}

func (e *IntoolsEngine) getGroups() []string {
    len := e.RedisClient.LLen("intools:groups").Val()
	return e.RedisClient.LRange("intools:groups", 0, len).Val()
}

func (e *IntoolsEngine) getConnectors(group string) []string {
    key := fmt.Sprintf("intools:groups:%s:connectors", group)
    len := e.RedisClient.LLen(key).Val()
    return e.RedisClient.LRange(key, 0, len).Val()
}