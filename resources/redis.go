package resources

const RedisResourceName = "redis"

type Redis struct {
	Name `yaml:"resource_name" env:"-"`

	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`

	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`
	Db   int    `yaml:"db"`
}

func NewRedis(n Name) Resource {
	return &Redis{
		Name: n,
		Host: "0.0.0.0",
		Port: 6379,
	}
}

func (p *Redis) GetType() string {
	return RedisResourceName
}
