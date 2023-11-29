package resources

const TelegramResourceName = "telegram"

const (
	TelegramApiKey = "TELEGRAM_API_KEY"
)

type Telegram struct {
	Name `yaml:"resource_name"`

	ApiKey string `yaml:"api_key"`
}

func (t *Telegram) GetType() string {
	return TelegramResourceName
}

func (t *Telegram) ToEnv() map[string]string {
	return map[string]string{
		TelegramApiKey:  t.ApiKey,
		EnvResourceName: t.GetName(),
	}
}

func (t *Telegram) FromEnv(in map[string]string) (err error) {
	t.ApiKey = in[TelegramApiKey]

	t.Name = Name(in[EnvResourceName])
	if t.Name == "" {
		t.Name = TelegramResourceName
	}

	return nil
}
