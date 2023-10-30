package resources

type Telegram struct {
	Name   string `yaml:"name"`
	ApiKey string `yaml:"api_key"`
}

func (t *Telegram) GetName() string {
	return t.Name
}

func (t *Telegram) setName(name string) {
	t.Name = name
}
