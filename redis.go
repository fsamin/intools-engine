package main

import "fmt"

func GetRedisGroupsKey() string {
	return "intools:groups"
}

func GetRedisGroupKey(group string) string {
    return "intools:groups:" + group
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
    multi := e.RedisClient.Multi()
    multi.Exec(func() error {
        multi.LPush(GetRedisGroupsKey(), c.Group)
        multi.LPush(c.GetRedisConnectorsKey(), c.Name)
        multi.Set(c.GetRedisConnectorKey(), c.GetJSON(), 0)
        return nil
    })
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
	len := e.RedisClient.LLen(GetRedisGroupsKey()).Val()
	return e.RedisClient.LRange(GetRedisGroupsKey(), 0, len).Val()
}

func (e *IntoolsEngine) getConnectors(group string) []string {
	key := fmt.Sprintf("intools:groups:%s:connectors", group)
	len := e.RedisClient.LLen(key).Val()
	return e.RedisClient.LRange(key, 0, len).Val()
}

func (e *IntoolsEngine) createGroup(group string) (bool, error) {
	listGroup := e.getGroups()
	for _, g := range listGroup {
		if group == g {
            return false, nil
		}
	}
    cmd1 := e.RedisClient.LPush(GetRedisGroupsKey(), group)
    if cmd1.Err() != nil {
        return false, cmd1.Err()
    } else {
        return true, nil
    }
}

func (e *IntoolsEngine) deleteGroup(group string) error {
	keyGroup := fmt.Sprintf("intools:groups:%s", group)
	cmd := e.RedisClient.Del(keyGroup)
    Debug.Printf("%s", cmd.Val())
	if cmd.Err() != nil {
		return cmd.Err()
	}
    //TODO delete branch
	key := GetRedisGroupsKey()
	cmd = e.RedisClient.LRem(key, 0, group)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}
