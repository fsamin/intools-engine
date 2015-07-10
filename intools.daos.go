package main

import "fmt"

func (e *IntoolsEngine) getGroups() []string {
	ret, err := RedisGetGroups(&e.RedisClient)
	if err != nil {
		Error.Printf("Error while getting groups from Redis %s", err.Error())
		return nil
	}
	return ret
}

func (e *IntoolsEngine) createGroup(group string) (bool, error) {
	return RedisCreateGroup(&e.RedisClient, group)
}

func (e *IntoolsEngine) saveConnector(c *Connector) {
	err := RedisSaveConnector(&e.RedisClient, c)
	if err != nil {
		Error.Printf("Error while saving to Redis %s", err.Error())
	}
}

func (e *IntoolsEngine) saveExecutor(c *Connector, exec *Executor) {
	err := RedisSaveExecutor(&e.RedisClient, c, exec)
	if err != nil {
		Error.Printf("Error while saving to Redis %s", err.Error())
	}
}

func (e *IntoolsEngine) getConnectors(group string) []string {
	ret, err := RedisGetConnectors(&e.RedisClient, group)
	if err != nil {
		Error.Printf("Error while getting connectors for group %s from Redis %s", group, err.Error())
		return nil
	}
	return ret
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
