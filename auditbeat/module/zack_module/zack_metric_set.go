package zack_module

import "github.com/elastic/beats/v7/metricbeat/mb"

type ZackMetricSet struct {
	mb.BaseMetricSet
	module *ZackModule
}

func NewZackMetricSet(base mb.BaseMetricSet) ZackMetricSet {
	return ZackMetricSet{
		BaseMetricSet: base,
		module:        base.Module().(*ZackModule),
	}
}

func (ms *ZackMetricSet) Hostname() string {
	return ms.module.hostName
}
