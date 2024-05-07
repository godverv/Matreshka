package grpc

import (
	"context"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) PatchConfigRaw(ctx context.Context, patch *matreshka_api.PatchConfigRaw_Request) (*matreshka_api.PatchConfigRaw_Response, error) {
	return &matreshka_api.PatchConfigRaw_Response{}, a.storage.PatchConfig(patch.ServiceName, patch.Raw)
}
