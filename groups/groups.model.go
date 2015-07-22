package groups

import (
	"github.com/soprasteria/intools-engine/connectors"
)

type Group struct {
	Name       string                 `json:"name"`
	Connectors []connectors.Connector `json:"connectors,omitempty"`
}
