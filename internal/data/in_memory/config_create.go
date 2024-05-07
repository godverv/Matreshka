package in_memory

import (
	"context"
	"time"

	"github.com/godverv/matreshka"
	"github.com/godverv/matreshka/api"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (d *data) CreateConfig(_ context.Context, req *matreshka_api.UpdateServiceConfig_Request) error {
	cfg := matreshka.NewEmptyConfig()
	c := &config{
		cfg: &cfg,
	}

	{
		appConfig := req.GetConfig().GetAppConfig()

		c.cfg.AppInfo = matreshka.AppInfo{
			Name:    appConfig.Name,
			Version: appConfig.Version,
		}

		c.cfg.StartupDuration, _ = time.ParseDuration(appConfig.StartupDuration)
	}

	{
		servers := req.GetConfig().GetApi()
		c.cfg.Servers = make(matreshka.Servers, len(servers))

		for i, item := range servers {
			switch item.ApiType {
			case matreshka_api.Config_Api_grpc:
				c.cfg.Servers[i] = &api.GRPC{
					Name: api.Name(item.MakoshName),
				}
			case matreshka_api.Config_Api_rest:
				c.cfg.Servers[i] = &api.Rest{
					Name: api.Name(item.MakoshName),
				}
			default:
				c.cfg.Servers[i] = &api.Unknown{
					Name: api.Name(item.MakoshName),
				}
			}
		}
	}

	{
		env := req.GetConfig().GetEnvironment()
		c.cfg.Environment = make(map[string]interface{}, len(env))
		for _, item := range env {
			c.cfg.Environment[item.Key] = item.Value
		}
	}

	d.m.Lock()
	d.mp[c.cfg.Name] = c
	d.m.Unlock()

	return nil
}
