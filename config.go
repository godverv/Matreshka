package matreshka

import (
	"time"

	"github.com/godverv/matreshka/resources"
	"github.com/godverv/matreshka/server"
)

type AppConfig struct {
	AppInfo     AppInfo              `yaml:"app_info"`
	DataSources []resources.Resource `yaml:"data_sources,omitempty"`
	Server      []server.Server      `yaml:"server,omitempty"`
}

type AppInfo struct {
	Name            string        `yaml:"name"`
	Version         string        `yaml:"version"`
	StartupDuration time.Duration `yaml:"startup_duration"`
}
