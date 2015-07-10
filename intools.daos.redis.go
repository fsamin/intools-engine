package main

import "gopkg.in/redis.v3"
import "fmt"

func GetRedisGroupsKey() string {
	return "intools:groups"
}

func GetRedisGroupKey(group string) string {
	return "intools:groups:" + group
}

func GetRedisConnectorsKey(c *Connector) string {
	return "intools:groups:" + c.Group + ":connectors"
}

func GetRedisConnectorKey(c *Connector) string {
	return "intools:groups:" + c.Group + ":connectors:" + c.Name
}

func GetRedisExecutorKey(c *Connector, e *Executor) string {
	return "intools:groups:" + c.Group + ":connectors:" + c.Name + ":executors"
}

func GetRedisResultKey(c *Connector, e *Executor) string {
	return "intools:groups:" + c.Group + ":connectors:" + c.Name + ":results"
}

func RedisSaveConnector(r *redis.Client, c *Connector) error {
	Debug.Printf("Saving %s to redis", c.Group)
	multi := r.Multi()
	_, err := multi.Exec(func() error {
		multi.LPush(GetRedisGroupsKey(), c.Group)
		multi.LPush(GetRedisConnectorsKey(c), c.Name)
		multi.Set(GetRedisConnectorKey(c), c.GetJSON(), 0)
		return nil
	})
	return err
}

func RedisSaveExecutor(r *redis.Client, c *Connector, exec *Executor) error {
	Debug.Printf("Saving %s:%s to redis", c.GetContainerName(), exec.ContainerId)
	cmd := r.LPush(GetRedisExecutorKey(c, exec), exec.GetJSON())
	if exec.Valid {
		_ = r.LPush(GetRedisResultKey(c, exec), exec.GetResult())
	}
	return cmd.Err()
}

func RedisGetGroups(r *redis.Client) ([]string, error) {
	len, err := r.LLen(GetRedisGroupsKey()).Result()
	if err != nil {
		return nil, err
	}
	return r.LRange(GetRedisGroupsKey(), 0, len).Result()
}

func RedisGetConnectors(r *redis.Client, group string) ([]string, error) {
	key := fmt.Sprintf("intools:groups:%s:connectors", group)
	len, err := r.LLen(key).Result()
	if err != nil {
		return nil, err
	}
	return r.LRange(key, 0, len).Result()
}

func RedisCreateGroup(r *redis.Client, group string) (bool, error) {
	listGroup, err := RedisGetGroups(r)
	if err != nil {
		return false, err
	}
	for _, g := range listGroup {
		if group == g {
			return false, nil
		}
	}
	cmd1 := r.LPush(GetRedisGroupsKey(), group)
	if cmd1.Err() != nil {
		return false, cmd1.Err()
	} else {
		return true, nil
	}
}
