package groups

import (
	"github.com/fsamin/intools-engine/common/logs"
	"github.com/fsamin/intools-engine/connectors"
)

func GetGroup(name string, withConnectors bool) *Group {
	allGroups := GetGroups(withConnectors)
	for _, g := range allGroups {
		if g.Name == name {
			return &g
		}
	}
	return nil
}

func GetGroups(withConnectors bool) []Group {
	groups, err := RedisGetGroups()
	if err != nil {
		logs.Error.Printf("Error while getting groups from Redis %s", err.Error())
		return nil
	}
	allGroups := make([]Group, len(groups))
	for i, g := range groups {
		group := &Group{
			Name: g,
		}
		if withConnectors {
			connectors := connectors.GetConnectors(g)
			group.Connectors = connectors
		}
		allGroups[i] = *group
	}
	return allGroups
}

func CreateGroup(group string) (bool, error) {
	return RedisCreateGroup(group)
}

func DeleteGroup(group string) error {
	return RedisDeleteGroup(group)
}
