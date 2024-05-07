package in_memory

import (
	"sync"

	"github.com/godverv/matreshka"
)

type config struct {
	m   sync.RWMutex
	cfg *matreshka.AppConfig
}

type data struct {
	m  sync.RWMutex
	mp map[string]*config
}

func New() *data {
	return &data{
		m:  sync.RWMutex{},
		mp: make(map[string]*config),
	}
}
