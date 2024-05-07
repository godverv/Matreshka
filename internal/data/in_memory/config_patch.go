package in_memory

import (
	"context"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (d *data) PatchEnvConfig(ctx context.Context, patch *matreshka_api.PatchConfigEnv_Request) error {
	d.m.Lock()

	c, ok := d.mp[patch.ServiceName]
	if ok {
		for _, p := range patch.Patches {
			c.cfg.Environment[p.FieldName] = p.Value
		}
	}

	d.m.Unlock()

	return nil
}
