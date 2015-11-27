package connectors

import (
	"encoding/json"
	"fmt"
	"github.com/fsamin/intools-engine/common/logs"
	"github.com/fsamin/intools-engine/executors"
	"github.com/fsamin/intools-engine/intools"
	"errors"
)

func GetRedisConnectorsKey(c *Connector) string {
	return "intools:groups:" + c.Group + ":connectors"
}

func GetRedisConnectorKey(c *Connector) string {
	return GetRedisrKey(c.Group, c.Name)
}

func GetRedisrKey(g string, c string) string {
	return "intools:groups:" + g + ":connectors:" + c
}

func GetRedisConnectorConfKey(g string, c string) string {
    return GetRedisrKey(g, c) + ":conf"
}

func RedisGetConnectors(group string) ([]string, error) {
	r := intools.Engine.GetRedisClient()
	key := fmt.Sprintf("intools:groups:%s:connectors", group)
	len, err := r.LLen(key).Result()
	if err != nil {
		return nil, err
	}
	return r.LRange(key, 0, len).Result()
}

func RedisGetConnector(group string, connector string) (*Connector, error) {
    r := intools.Engine.GetRedisClient()
	logs.Debug.Printf("Loading %s:%s from redis", group, connector)
	key := GetRedisConnectorConfKey(group, connector)
	cmd := r.Get(key)
	jsonCmd := cmd.Val()
	if cmd.Err() != nil {
		logs.Error.Printf("Redis command failed %s", cmd.Err())
		return nil, errors.New("Unable to load connectors " + group + "/" + connector + ":" + cmd.Err().Error())
	}
	c := &Connector{}
	err := json.Unmarshal([]byte(jsonCmd), c)
	if err != nil {
		logs.Error.Printf("JSON Unmarshall failed with following value")
		logs.Error.Print(jsonCmd)
		return nil, err
	}
	return c, nil
}

func RedisSaveConnector(c *Connector) error {
	r := intools.Engine.GetRedisClient()
	logs.Debug.Printf("Saving %s to redis", c.Group)
	multi := r.Multi()
	_, err := multi.Exec(func() error {
		multi.LRem(GetRedisrKey(c.Group, c.Name), 0, c.Group)
		multi.LPush(GetRedisrKey(c.Group, c.Name), c.Group)
		multi.LRem(GetRedisConnectorsKey(c), 0, c.Name)
		multi.LPush(GetRedisConnectorsKey(c), c.Name)
		multi.Set(GetRedisConnectorConfKey(c.Group, c.Name), c.GetJSON(), 0)
		return nil
	})
	return err
}

func GetRedisExecutorKey(c *Connector) string {
	return "intools:groups:" + c.Group + ":connectors:" + c.Name + ":executors"
}

func GetRedisResultKey(c *Connector, e *executors.Executor) string {
	return "intools:groups:" + c.Group + ":connectors:" + c.Name + ":results"
}

func RedisSaveExecutor(c *Connector, exec *executors.Executor) error {
    r := intools.Engine.GetRedisClient()
    logs.Debug.Printf("Saving %s:%s to redis", c.GetContainerName(), exec.ContainerId)
	cmd := r.LPush(GetRedisExecutorKey(c), exec.GetJSON())
	if exec.Valid {
		_ = r.LPush(GetRedisResultKey(c, exec), exec.GetResult())
	}
	return cmd.Err()
}

func RedisGetLastExecutor(c *Connector) (string, error) {
    r := intools.Engine.GetRedisClient()
    cmd := r.LIndex(GetRedisExecutorKey(c), 0)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return cmd.Val(), nil
}
