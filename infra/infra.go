package infra

import (
	"ddd_demo/infra/consul"

	"github.com/google/wire"
)

var InfraSet = wire.NewSet(consul.NewConsulService)
