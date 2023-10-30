package env_parser

type EnvParser interface {
	// ToEnv - returns a set of NAME-VALUE environment variables required for this resource to be run
	ToEnv() map[string]string
	// FromEnv - fills fields from environment
	FromEnv(in map[string]string) (err error)
}
