package matreshka

import (
	"time"
)

type AppInfo struct {
	Name            string        `yaml:"name,omitempty" env:",omitempty"`
	Version         string        `yaml:"version,omitempty" env:",omitempty"`
	StartupDuration time.Duration `yaml:"startup_duration,omitempty" env:",omitempty"`
}
