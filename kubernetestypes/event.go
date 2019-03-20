package kubernetestypes

import (
	"strings"
	"time"
)

type Event struct {
	StageTimestamp time.Time
	Level          string
	Stage          string
	RequestURI     string
	Verb           string
	User           User
	SourceIPs      []string
	ObjectRef      ObjectRef
}

func (e *Event) GetSourceIPAddress() string {
	if len(e.SourceIPs) == 0 {
		return "UNKNOWN"
	} else if len(e.SourceIPs) == 1 {
		return e.SourceIPs[0]
	} else {
		return strings.Join(e.SourceIPs, ",")
	}
}

// TODO: JUSTIN: Figure out how to use the json attributes on the field names (so I can call the field name "Timestamp" and have it load the JSON field "stageTimestamp")

// TODO: JUSTIN: Get the following object (annotations) from the JSON post into this object
// annotations
// authorization.k8s.io/decision
// authorization.k8s.io/reason
