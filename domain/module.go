package domain

import (
	"learnfx/domain/service"

	"go.uber.org/fx"
)

var Module = fx.Module("domain",
	fx.Provide(service.NewItemDomainServiceImpl),
)
