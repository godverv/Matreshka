package matreshka

import (
	"time"
)

type AppInfo struct {
	Name            string        `yaml:"name"`
	Version         string        `yaml:"version"`
	StartupDuration time.Duration `yaml:"startup_duration"`
}
