package grpc

import (
	"context"

	api "github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) UpdateServiceConfig(ctx context.Context, req *api.UpdateServiceConfig_Request) (*api.UpdateServiceConfig_Response, error) {
	return &api.UpdateServiceConfig_Response{}, a.storage.CreateConfig(ctx, req)
}
