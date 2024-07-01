package matreshka

import (
	"time"
)

type AppInfo struct {
	Name            string        `yaml:"name" env:",omitempty"`
	Version         string        `yaml:"version" env:",omitempty"`
	StartupDuration time.Duration `yaml:"startup_duration" env:",omitempty"`
}
