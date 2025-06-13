package user

import (
	"net/http"

	types "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/types"
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	controllers "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/controllers"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/usecases"

	"go.uber.org/fx"
)

func configureModuleRoutes(
	ctrlGetAllUsers *controllers.GetAllUsersController,
	h *types.HandlersStore,
) {

	handlersModuleUsers := &types.SliceHandlers{
		Prefix: "users",
		Routes: []types.HandlerModule{
			{
				Route:   "/",
				Method:  http.MethodGet,
				Handler: ctrlGetAllUsers.Run,
			},
		},
	}

	h.Handlers = append(h.Handlers, *handlersModuleUsers)
}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(dao.NewMySQLUserDao),
		fx.Provide(usecases.NewGetAllUsers),
		fx.Provide(controllers.NewGetAllUsersController),

		fx.Invoke(configureModuleRoutes),
	}
}
