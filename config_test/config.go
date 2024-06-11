package config

import (
	"time"
)

type EnvironmentConfig struct {
	AvailablePorts                   []int
	CreditPercent                    float64
	CreditPercentsBasedOnYearOfBirth []float64
	DatabaseMaxConnections           int
	OneOfWelcomeString               string
	RequestTimeout                   time.Duration
	TrueFalser                       bool
	UsernamesToBan                   []string
	WelcomeString                    string
}
