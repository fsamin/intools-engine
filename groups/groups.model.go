package groups

import (
	"github.com/fsamin/intools-engine/connectors"
)

type Group struct {
	Name       string                 `json:"name"`
	Connectors []connectors.Connector `json:"connectors,omitempty"`
}
