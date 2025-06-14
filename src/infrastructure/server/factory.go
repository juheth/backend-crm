package infraestructure

import (
	config "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/config"
	result "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	types "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/types"
	db "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/adapter"

	"go.uber.org/fx"
)

type ProvidersStore struct {
	Providers []fx.Option
}

func (ps *ProvidersStore) Init() {
	ps.Providers = []fx.Option{
		fx.Provide(types.NewHandlersStore),
		fx.Provide(result.NewResult),
		fx.Provide(config.NewConfig),
		fx.Provide(db.NewDBConnection),
	}
}
func (ps *ProvidersStore) AddModule(p []fx.Option) {
	ps.Providers = append(ps.Providers, p...)
}

func (ps *ProvidersStore) Up(lp ...[]fx.Option) {
	ps.Providers = append(ps.Providers, fx.Invoke(NewHttpFiberServer))
	fx.New(ps.Providers...).Run()
}
