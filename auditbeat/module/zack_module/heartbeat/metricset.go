package heartbeat

import (
	"github.com/elastic/beats/v7/auditbeat/module/zack_module"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/mb/parse"
	"github.com/elastic/elastic-agent-libs/logp"
)

const (
	moduleName    = "zack"
	metricSetName = "heartbeat"
	namespace     = "."
)

type MetricSet struct {
	zack_module.ZackMetricSet
	// We can add config or other attibutes implementation here for HeartBeat MetricBeat
	log *logp.Logger
}

func init() {
	mb.Registry.MustAddMetricSet(moduleName, metricSetName, New,
		mb.DefaultMetricSet(), mb.WithHostParser(parse.EmptyHostParser), mb.WithNamespace(namespace))
}

func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	return &MetricSet{
		ZackMetricSet: zack_module.NewZackMetricSet(base),
		log:           logp.NewLogger(metricSetName),
	}, nil
}

func (ms *MetricSet) Fetch(reporter mb.ReporterV2) {
	event := HeartBeatEvent{
		HostName: ms.ZackMetricSet.Hostname(),
	}
	ms.log.Info("sending heartbeat event")
	reporter.Event(event.buildHeartBeatEvent())
}
