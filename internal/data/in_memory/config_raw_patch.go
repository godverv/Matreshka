package in_memory

import (
	"github.com/godverv/matreshka"
	"github.com/pkg/errors"
)

func (d *data) PatchConfig(name string, raw []byte) error {
	var cfg matreshka.AppConfig

	err := cfg.Unmarshal(raw)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling")
	}

	d.m.Lock()
	d.mp[name] = &config{cfg: &cfg}
	d.m.Unlock()

	return nil
}
