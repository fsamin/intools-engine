package groups

import "github.com/fsamin/intools-engine/intools"
import "fmt"
import "github.com/fsamin/intools-engine/common/logs"

func GetRedisGroupsKey() string {
	return "intools:groups"
}

func GetRedisGroupKey(group string) string {
	return "intools:groups:" + group
}

func RedisGetLength() (int64, error) {
	r := intools.Engine.GetRedisClient()
	len, err := r.LLen(GetRedisGroupsKey()).Result()
	if err != nil {
		return 0, err
	}
	return len, nil
}

func RedisGetGroups() ([]string, error) {
	r := intools.Engine.GetRedisClient()
	len, err := r.LLen(GetRedisGroupsKey()).Result()
	if err != nil {
		return nil, err
	}
	return r.LRange(GetRedisGroupsKey(), 0, len).Result()
}

func RedisCreateGroup(group string) (bool, error) {
	r := intools.Engine.GetRedisClient()
	listGroup, err := RedisGetGroups()
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

func RedisDeleteGroup(group string) error {
	r := intools.Engine.GetRedisClient()
	keyGroup := fmt.Sprintf("intools:groups:%s:*", group)
	evalKeys := make([]string, 0)
	evalArgs := []string{keyGroup}
	//TODO This remove too much data
	cmd := r.Eval("return redis.call('del', unpack(redis.call('keys', ARGV[1])))", evalKeys, evalArgs)

	logs.Debug.Printf("%s --> %d", keyGroup, cmd.Val())
	if cmd.Err() != nil {
		return cmd.Err()
	}
	//TODO delete branch
	key := GetRedisGroupsKey()
	cmd1 := r.LRem(key, 0, group)
	if cmd1.Err() != nil {
		return cmd1.Err()
	}
	return nil
}
