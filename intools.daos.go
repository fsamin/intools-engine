package main

import "fmt"
import "encoding/json"

func (e *IntoolsEngine) GetGroup(name string, withConnectors bool) *Group {
	allGroups := e.GetGroups(withConnectors)
	for _, g := range allGroups {
		if g.Name == name {
			return &g
		}
	}
	return nil
}

func (e *IntoolsEngine) GetGroups(withConnectors bool) []Group {
	groups, err := RedisGetGroups(&e.RedisClient)
	if err != nil {
		Error.Printf("Error while getting groups from Redis %s", err.Error())
		return nil
	}
	allGroups := make([]Group, len(groups))
	for i, g := range groups {
		group := &Group{
			Name: g,
		}
		if withConnectors {
			connectors := e.GetConnectors(g)
			group.Connectors = connectors
		}
		allGroups[i] = *group
	}
	return allGroups
}

func (e *IntoolsEngine) CreateGroup(group string) (bool, error) {
	return RedisCreateGroup(&e.RedisClient, group)
}

func (e *IntoolsEngine) SaveConnector(c *Connector) {
	err := RedisSaveConnector(&e.RedisClient, c)
	if err != nil {
		Error.Printf("Error while saving to Redis %s", err.Error())
	}
}

func (e *IntoolsEngine) SaveExecutor(c *Connector, exec *Executor) {
	err := RedisSaveExecutor(&e.RedisClient, c, exec)
	if err != nil {
		Error.Printf("Error while saving to Redis %s", err.Error())
	}
}

func (e *IntoolsEngine) GetLastConnectorExecutor(c *Connector) *Executor {
	sExecutor, err := RedisGetLastExecutor(&e.RedisClient, c)
	if err != nil {
		Error.Printf("Cannot load last executor %s:%s from Redis", c.Group, c.Name)
		return nil
	}
	executor := &Executor{}
	err = json.Unmarshal([]byte(sExecutor), executor)
	if err != nil {
		Error.Printf("Cannot parse last executor %s:%s", c.Group, c.Name)
		Error.Printf(err.Error())
		return nil
	}
	return executor
}

func (e *IntoolsEngine) GetConnector(group string, connector string) (*Connector, error) {
	conn, err := RedisGetConnector(&e.RedisClient, group, connector)
	if err != nil {
		Error.Printf("Error while loading %s:%s to Redis %s", group, connector, err.Error())
		return nil, err
	}
	return conn, nil
}

func (e *IntoolsEngine) GetConnectors(group string) []Connector {
	ret, err := RedisGetConnectors(&e.RedisClient, group)
	if err != nil {
		Error.Printf("Error while getting connectors for group %s from Redis %s", group, err.Error())
		return nil
	}
	connectors := make([]Connector, len(ret))
	for i, c := range ret {
		conn, err := e.GetConnector(group, c)
		if err != nil {
			Warning.Printf("Unable to load %s:%s : %s", group, c, err)
		} else {
			connectors[i] = *conn
		}
	}
	return connectors
}

func (e *IntoolsEngine) DeleteGroup(group string) error {
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
