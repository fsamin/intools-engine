package executors

import (
	"encoding/json"
	"github.com/fsamin/intools-engine/common/logs"
	"time"
)

type Executor struct {
	ContainerId string
	Host        string
	Running     bool
	Terminated  bool
	ExitCode    int
	Stdout      string
	JsonStdout  *map[string]interface{}
	Stderr      string
	StartedAt   time.Time
	FinishedAt  time.Time
	Valid       bool
}

func (e *Executor) GetJSON() string {
	b, err := json.Marshal(e)
	if err != nil {
		logs.Error.Println(err)
		return ""
	}
	return string(b[:])
}

func (e *Executor) GetResult() string {
	b, err := json.Marshal(e.JsonStdout)
	if err != nil {
		logs.Error.Println(err)
		return ""
	}
	return string(b[:])
}
