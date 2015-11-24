package connectors

import (
	"encoding/json"
	"github.com/fsamin/intools-engine/common/logs"
	"github.com/fsamin/intools-engine/executors"
)

func SaveExecutor(c *Connector, exec *executors.Executor) {
	err := RedisSaveExecutor(c, exec)
	if err != nil {
		logs.Error.Printf("Error while saving to Redis %s", err.Error())
	}
}

func SaveConnector(c *Connector) {
	err := RedisSaveConnector(c)
	if err != nil {
		logs.Error.Printf("Error while saving to Redis %s", err.Error())
	}
}

func GetLastConnectorExecutor(c *Connector) *executors.Executor {
	sExecutor, err := RedisGetLastExecutor(c)
	if err != nil {
		logs.Error.Printf("Cannot load last executor %s:%s from Redis", c.Group, c.Name)
		return nil
	}
	executor := &executors.Executor{}
	err = json.Unmarshal([]byte(sExecutor), executor)
	if err != nil {
		logs.Error.Printf("Cannot parse last executor %s:%s", c.Group, c.Name)
		logs.Error.Printf(err.Error())
		return nil
	}
	return executor
}

func GetConnector(group string, connector string) (*Connector, error) {
	conn, err := RedisGetConnector(group, connector)
	if err != nil {
		logs.Error.Printf("Error while loading %s:%s to Redis %s", group, connector, err.Error())
		return nil, err
	}
	return conn, nil
}

func GetConnectors(group string) []Connector {
	ret, err := RedisGetConnectors(group)
	if err != nil {
		logs.Error.Printf("Error while getting connectors for group %s from Redis %s", group, err.Error())
		return nil
	}
	connectors := make([]Connector, len(ret))
	for i, c := range ret {
		conn, err := GetConnector(group, c)
		if err != nil {
			logs.Warning.Printf("Unable to load %s:%s : %s", group, c, err)
		} else {
			connectors[i] = *conn
		}
	}
	return connectors
}
