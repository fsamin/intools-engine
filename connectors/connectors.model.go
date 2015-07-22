package connectors

import (
	"encoding/json"
	"github.com/samalba/dockerclient"
	"github.com/soprasteria/intools-engine/common/logs"
	"github.com/soprasteria/intools-engine/executors"
)

type Connector struct {
	Group           string                        `json:"group"`
	Name            string                        `json:"name"`
	ContainerConfig *dockerclient.ContainerConfig `json:"config"`
	Timeout         int                           `json:"timeout,omitempty"`
	Refresh         int                           `json:"refresh,omitempty"`
}

type ConnectorRunner interface {
	Exec(connector *Connector) (*executors.Executor, error)
}

func NewConnector(group string, name string) *Connector {
	conn := &Connector{group, name, nil, 15, 60}
	return conn
}

func (c *Connector) Init(image string, timeout int, refresh int, cmd []string) {
	if c.ContainerConfig == nil {
		c.ContainerConfig = &dockerclient.ContainerConfig{
			Image:        image,
			Cmd:          cmd,
			AttachStdin:  false,
			AttachStdout: false,
			AttachStderr: false,
			Tty:          false,
		}
	}

	if timeout != 0 {
		c.Timeout = timeout
	}
	if refresh != 0 {
		c.Refresh = refresh
	}
}

func (c *Connector) GetContainerName() string {
	return c.Group + "-" + c.Name
}

func (c *Connector) GetJSON() string {
	b, err := json.Marshal(c)
	if err != nil {
		logs.Error.Println(err)
		return ""
	}
	return string(b[:])
}

func (c *Connector) Run() {
	//TODO : Should not run error, or invalid connector ?
	logs.Debug.Printf("Run Connector %s:%s", c.Group, c.Name)
	Exec(c)
}
