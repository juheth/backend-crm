package user

import (
	"net/http"

	types "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/types"
	controllers "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/controllers"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/usecases"

	"go.uber.org/fx"
)

func configureModuleRoutes(
	ctrlFindAllUsers *controllers.FindAllUserController,
	h *types.HandlersStore,
) {

	handlersModuleUsers := &types.SliceHandlers{
		Prefix: "users",
		Routes: []types.HandlerModule{
			{
				Route:   "/",
				Method:  http.MethodGet,
				Handler: ctrlFindAllUsers.Run,
			},
		},
	}

	h.Handlers = append(h.Handlers, *handlersModuleUsers)
}

func ModuleProviders() []fx.Option {
	return []fx.Option{

		fx.Provide(usecases.NewUserUseCase),
		fx.Provide(controllers.NewFindAllUserController),

		fx.Invoke(configureModuleRoutes),
	}
}
