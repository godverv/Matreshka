package grpc

import (
	"context"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) GetConfig(ctx context.Context, request *matreshka_api.GetConfig_Request) (*matreshka_api.GetConfig_Response, error) {
	return a.storage.GetConfig(ctx, request)
}
