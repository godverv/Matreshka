package matreshka

import (
	"time"
)

type AppInfo struct {
	Name            string        `yaml:"name" env:"name"`
	Version         string        `yaml:"version" env:"version"`
	StartupDuration time.Duration `yaml:"startup_duration" env:"startup_duration"`
}
