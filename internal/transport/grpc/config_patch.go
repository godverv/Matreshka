package grpc

import (
	"context"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) PatchConfigEnv(ctx context.Context, patch *matreshka_api.PatchConfigEnv_Request) (*matreshka_api.PatchConfigEnv_Response, error) {
	return &matreshka_api.PatchConfigEnv_Response{}, a.storage.PatchEnvConfig(ctx, patch)
}
