package heartbeat

import (
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

type HeartBeatEvent struct {
	HostName string
}

func (e *HeartBeatEvent) buildHeartBeatEvent() mb.Event {
	data := mapstr.M{
		"host_name": e.HostName,
	}
	return mb.Event{
		RootFields: mapstr.M{
			"heartbeat": data,
		},
	}
}
