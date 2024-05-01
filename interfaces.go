package matreshka

import (
	"time"

	"github.com/godverv/matreshka/api"
	"github.com/godverv/matreshka/resources"
)

type Config interface {
	AppInfo() AppInfo
	Api() API
	Resources() Resource

	GetInt(key string) (out int)
	GetString(key string) (out string)
	GetBool(key string) (out bool)
	GetDuration(key string) (out time.Duration)

	TryGetInt(key string) (out int, err error)
	TryGetString(key string) (out string, err error)
	TryGetBool(key string) (out bool, err error)
	TryGetDuration(key string) (t time.Duration, err error)

	GetMatreshka() *AppConfig
}

type API interface {
	REST(name string) (*api.Rest, error)
	GRPC(name string) (*api.GRPC, error)
}

type Resource interface {
	Postgres(name string) (*resources.Postgres, error)
	Telegram(name string) (*resources.Telegram, error)
	Redis(name string) (*resources.Redis, error)
	GRPC(name string) (*resources.GRPC, error)
	Sqlite(name string) (*resources.Sqlite, error)
}
