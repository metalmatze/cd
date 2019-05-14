package signalcd

import (
	"encoding/json"
	"time"
)

type PipelineStatus string

const (
	Unknown  PipelineStatus = "unknown"
	Success  PipelineStatus = "success"
	Failed   PipelineStatus = "failed"
	Progress PipelineStatus = "progress"
)

type Agent struct {
	Name     string         `json:"name"`
	Status   PipelineStatus `json:"status"`
	Pipeline Pipeline       `json:"pipeline"`
}

type AgentServer struct {
	Agent `json:",inline"`

	Heartbeat time.Time `json:"-"`
}

// Ready returns if an Agent is ready.
// If ready an agent has reported to the API in the past 15s.
func (as AgentServer) Ready() bool {
	return time.Since(as.Heartbeat) < 15*time.Second
}

func (as AgentServer) MarshalJSON() ([]byte, error) {
	s := struct {
		Agent `json:",inline"`
		Ready bool `json:"ready"`
	}{
		Agent: as.Agent,
		Ready: as.Ready(),
	}

	return json.Marshal(s)
}
