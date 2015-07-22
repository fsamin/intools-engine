package groups

import (
	"github.com/soprasteria/intools-engine/common/logs"
	"github.com/soprasteria/intools-engine/connectors"
)

func Reload() {
	groups := GetGroups(true)
	for _, group := range groups {
		logs.Trace.Printf("%s - Reloading group", group.Name)
		for _, connector := range group.Connectors {
			logs.Trace.Printf("%s:%s - Reloading connector", group.Name, connector.Name)
			connectors.InitSchedule(&connector)
		}
	}
}
