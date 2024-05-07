package in_memory

import (
	"context"
	"fmt"

	"github.com/godverv/matreshka"
	"github.com/godverv/matreshka/resources"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

var resourceToProtoEnum = map[string]matreshka_api.Config_Resource_Type{
	resources.PostgresResourceName: matreshka_api.Config_Resource_postgres,
	resources.RedisResourceName:    matreshka_api.Config_Resource_redis,
	resources.TelegramResourceName: matreshka_api.Config_Resource_telegram,
	resources.GrpcResourceName:     matreshka_api.Config_Resource_grpc,
}

func (d *data) GetConfig(_ context.Context, config *matreshka_api.GetConfig_Request) (*matreshka_api.GetConfig_Response, error) {
	d.m.RLock()
	a, _ := d.mp[config.ServiceName]

	out := &matreshka_api.GetConfig_Response{}

	if a != nil {
		out.Config = cfgToProto(a.cfg)
	}

	d.m.RUnlock()

	return out, nil
}

func (d *data) GetRawConfig(_ context.Context, config *matreshka_api.GetConfigRaw_Request) (*matreshka_api.GetConfigRaw_Response, error) {
	d.m.RLock()
	a, _ := d.mp[config.ServiceName]
	d.m.RUnlock()

	out := &matreshka_api.GetConfigRaw_Response{}

	if a != nil {
		b, _ := a.cfg.Marshal()
		out.Config = b
	}

	return out, nil
}

func cfgToProto(cfg *matreshka.AppConfig) *matreshka_api.Config {
	if cfg == nil {
		return nil
	}

	out := &matreshka_api.Config{
		Resources:   make([]*matreshka_api.Config_Resource, len(cfg.Resources)),
		Environment: make([]*matreshka_api.Config_Environment, 0, len(cfg.Environment)),
	}

	out.AppConfig = &matreshka_api.Config_AppConfig{
		Name:            cfg.Name,
		Version:         cfg.Version,
		StartupDuration: cfg.StartupDuration.String(),
	}

	for i, r := range cfg.Resources {
		out.Resources[i] = &matreshka_api.Config_Resource{
			ResourceType:     resourceToProtoEnum[r.GetType()],
			ConnectionString: "TODO", // TODO
		}
	}

	for k, v := range cfg.Environment {
		out.Environment = append(out.Environment, &matreshka_api.Config_Environment{
			Key:   k,
			Value: fmt.Sprintf("%s", v),
		})
	}

	return out
}
