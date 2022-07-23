package zack_module

import (
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/go-sysinfo"
)

const moduleName = "zack"

func init() {
	if err := mb.Registry.AddModule(moduleName, NewModule); err != nil {
		panic(err)
	}
}

type ZackModule struct {
	mb.BaseModule
	hostName string
}

func NewModule(base mb.BaseModule) (mb.Module, error) {
	log := logp.NewLogger(moduleName)
	var hostName string
	if hostInfo, err := sysinfo.Host(); err != nil {
		log.Errorf("Could not get host info. err=%+v", err)
	} else {
		hostName = hostInfo.Info().Hostname
	}

	if hostName == "" {
		log.Warnf("Could not get host name")
	}

	return &ZackModule{
		BaseModule: base,
		hostName:   hostName,
	}, nil
}
