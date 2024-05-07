package data

import (
	"context"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

type Data interface {
	CreateConfig(ctx context.Context, config *matreshka_api.UpdateServiceConfig_Request) error
	GetConfig(ctx context.Context, config *matreshka_api.GetConfig_Request) (*matreshka_api.GetConfig_Response, error)
	GetRawConfig(ctx context.Context, config *matreshka_api.GetConfigRaw_Request) (*matreshka_api.GetConfigRaw_Response, error)
	PatchEnvConfig(ctx context.Context, patch *matreshka_api.PatchConfigEnv_Request) error
	PatchConfig(name string, raw []byte) error
}
