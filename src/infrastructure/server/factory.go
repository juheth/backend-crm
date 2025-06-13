package infrastructure

import (
	"github.com/juheth/Go-Clean-Arquitecture/src/common/auth"
	config "github.com/juheth/Go-Clean-Arquitecture/src/common/config"
	result "github.com/juheth/Go-Clean-Arquitecture/src/common/response"
	types "github.com/juheth/Go-Clean-Arquitecture/src/common/types"
	db "github.com/juheth/Go-Clean-Arquitecture/src/infrastructure/db/adapter"

	"github.com/juheth/Go-Clean-Arquitecture/src/modules/users"

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

		fx.Provide(func(cfg *config.Config) *auth.JWT {
			if cfg.App.SecretKey == "" {
				panic("JWT secret key no configurada")
			}
			return auth.NewJWT(cfg.App.SecretKey)
		}),
	}

	ps.AddModule(users.ModuleProviders())

}

func (ps *ProvidersStore) AddModule(p []fx.Option) {
	ps.Providers = append(ps.Providers, p...)
}

func (ps *ProvidersStore) Up(lp ...[]fx.Option) {
	ps.Providers = append(ps.Providers, fx.Invoke(NewHttpFiberServer))
	fx.New(ps.Providers...).Run()
}
