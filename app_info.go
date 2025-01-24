package matreshka

import (
	"strings"
	"time"
)

type AppInfo struct {
	Name            string        `yaml:"name,omitempty" env:",omitempty"`
	Version         string        `yaml:"version,omitempty" env:",omitempty"`
	StartupDuration time.Duration `yaml:"startup_duration,omitempty" env:",omitempty"`
}

func (a AppInfo) ModuleName() string {
	if len(a.Name) == 0 {
		return ""
	}

	startIdx := strings.LastIndex(a.Name, "/")
	startIdx++
	if startIdx != 0 {
		return a.Name[startIdx:]
	}

	return a.Name
}
