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
	ctrlGetUserById *controllers.GetUserByIdController,
	ctrlCrearUser *controllers.CreateUsersController,
	ctrlUpdateUser *controllers.UpdateUserController,

	h *types.HandlersStore,
) {

	handlersModuleUsers := &types.SliceHandlers{
		Prefix: "users",
		Routes: []types.HandlerModule{
			{
				Route:   "/all",
				Method:  http.MethodGet,
				Handler: ctrlGetAllUsers.Run,
			},
			{
				Route:   "/:id",
				Method:  http.MethodGet,
				Handler: ctrlGetUserById.Run,
			},
			{
				Route:   "/create",
				Method:  http.MethodPost,
				Handler: ctrlCrearUser.Run,
			},
			{
				Route:   "/update/:id",
				Method:  http.MethodPut,
				Handler: ctrlUpdateUser.Run,
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleUsers)
}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(dao.NewMySQLUserDao),
		fx.Provide(controllers.NewGetAllUsersController),
		fx.Provide(usecases.NewGetAllUsers),
		fx.Provide(controllers.NewGetUserByIdController),
		fx.Provide(usecases.NewGetUserById),
		fx.Provide(controllers.NewCreateUsersController),
		fx.Provide(usecases.NewCreateUsers),
		fx.Provide(controllers.NewUpdateUserController),
		fx.Provide(usecases.NewUpdateUser),

		fx.Invoke(configureModuleRoutes),
	}
}
