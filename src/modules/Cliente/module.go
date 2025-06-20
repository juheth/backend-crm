package client

import (
	"net/http"

	types "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/types"
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	controllers "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/controllers"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/usecases"

	"go.uber.org/fx"
)

func configureModuleRoutes(
	ctrlCreateClient *controllers.CreateClientController,
	h *types.HandlersStore,
) {

	handlersModuleClients := &types.SliceHandlers{
		Prefix: "clients",
		Routes: []types.HandlerModule{
			{
				Route:        "/create",
				Method:       http.MethodPost,
				Handler:      ctrlCreateClient.Run,
				RequiresAuth: false,
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleClients)
}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(dao.NewMySQLClientDao),
		fx.Provide(controllers.NewCreateClientController),
		fx.Provide(usecases.NewCreateClient),
		fx.Invoke(configureModuleRoutes),
	}
}
