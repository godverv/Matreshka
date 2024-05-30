package resources

const TelegramResourceName = "telegram"

type Telegram struct {
	Name `yaml:"resource_name" env:"-"`

	ApiKey string `yaml:"api_key"`
}

func NewTelegram(n Name) Resource {
	return &Telegram{
		Name: n,
	}
}

func (t *Telegram) GetType() string {
	return TelegramResourceName
}
