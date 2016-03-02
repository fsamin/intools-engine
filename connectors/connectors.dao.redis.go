package connectors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fsamin/intools-engine/common/logs"
	"github.com/fsamin/intools-engine/executors"
	"github.com/fsamin/intools-engine/intools"
	"gopkg.in/robfig/cron.v2"
	"strconv"
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

func GetRedisCronIdKey(g string, c string) string {
	return GetRedisrKey(g, c) + ":cronId"
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

func RedisGetConnectorCronId(group string, connector string) (cron.EntryID, error) {
	r := intools.Engine.GetRedisClient()
	logs.Debug.Printf("Fetching cronId for %s:%s from redis", group, connector)
	key := GetRedisCronIdKey(group, connector)
	cronId := r.Get(key)
	stCronId := cronId.Val()
	if cronId.Err() != nil {
		logs.Error.Printf("Redis command failed %s", cronId.Err())
		return 0, errors.New("Unable to fetch connectors cronId " + group + "/" + connector + ":" + cronId.Err().Error())
	}

	ret, err := strconv.Atoi(stCronId)
	if err != nil {
		logs.Error.Printf("Cannot convert cronId to interged")
		logs.Error.Print(stCronId)
		return 0, err
	}
	return cron.EntryID(ret), nil
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

func RedisRemoveConnector(c *Connector) error {
	r := intools.Engine.GetRedisClient()
	logs.Debug.Printf("Removing %s:%s from redis", c.Group, c.Name)
	multi := r.Multi()
	_, err := multi.Exec(func() error {
		multi.Del(GetRedisConnectorConfKey(c.Group, c.Name))
		multi.Del(GetRedisCronIdKey(c.Group, c.Name))
		multi.Del(GetRedisExecutorKey(c))
		multi.Del(GetRedisResultKey(c))
		multi.Del(GetRedisConnectorsKey(c))
		multi.Del(GetRedisConnectorsKey(c))
		multi.Del(GetRedisrKey(c.Group, c.Name))
		multi.Del(GetRedisrKey(c.Group, c.Name))
		return nil
	})
	return err
}

func GetRedisExecutorKey(c *Connector) string {
	return "intools:groups:" + c.Group + ":connectors:" + c.Name + ":executors"
}

func GetRedisResultKey(c *Connector) string {
	return "intools:groups:" + c.Group + ":connectors:" + c.Name + ":results"
}

func RedisSaveExecutor(c *Connector, exec *executors.Executor) error {
	r := intools.Engine.GetRedisClient()
	logs.Debug.Printf("Saving %s:%s to redis", c.GetContainerName(), exec.ContainerId)
	cmd := r.LPush(GetRedisExecutorKey(c), exec.GetJSON())
	if exec.Valid {
		_ = r.LPush(GetRedisResultKey(c), exec.GetResult())
	}
	return cmd.Err()
}

func RedisSaveCronId(c *Connector, cronId cron.EntryID) error {
	r := intools.Engine.GetRedisClient()
	logs.Debug.Printf("Saving cronId %s:%s to redis", c.GetContainerName(), cronId)
	cmd := r.Set(GetRedisCronIdKey(c.Group, c.Name), strconv.Itoa(int(cronId)), 0)
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
