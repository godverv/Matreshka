package grpc

import (
	"context"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) GetConfigRaw(ctx context.Context, req *matreshka_api.GetConfigRaw_Request) (*matreshka_api.GetConfigRaw_Response, error) {
	return a.storage.GetRawConfig(ctx, req)
}
