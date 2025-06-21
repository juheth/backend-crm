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
	ctrlGetAllClients *controllers.GetAllClientsController,
	ctrlGetClientByID *controllers.GetClientByIDController,
	ctrlUpdateClient *controllers.UpdateClientController,
	ctrlGetClientsByCreator *controllers.GetClientsByCreatorController,
	h *types.HandlersStore,
) {

	handlersModuleClients := &types.SliceHandlers{
		Prefix: "clients",
		Routes: []types.HandlerModule{
			{
				Route:        "/create",
				Method:       http.MethodPost,
				Handler:      ctrlCreateClient.Run,
				RequiresAuth: true,
			},
			{
				Route:        "/get-all",
				Method:       http.MethodGet,
				Handler:      ctrlGetAllClients.Run,
				RequiresAuth: true,
			},
			{
				Route:        "/get/:id",
				Method:       http.MethodGet,
				Handler:      ctrlGetClientByID.Run,
				RequiresAuth: true,
			},
			{
				Route:        "/update/:id",
				Method:       http.MethodPut,
				Handler:      ctrlUpdateClient.Run,
				RequiresAuth: true,
			},
			{
				Route:        "/creator/:creatorId",
				Method:       http.MethodGet,
				Handler:      ctrlGetClientsByCreator.Run,
				RequiresAuth: true,
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
		fx.Provide(controllers.NewGetAllClientsController),
		fx.Provide(usecases.NewGetAllClients),
		fx.Provide(controllers.NewGetClientByIDController),
		fx.Provide(usecases.NewGetClientById),
		fx.Provide(controllers.NewUpdateClientController),
		fx.Provide(usecases.NewUpdateClient),
		fx.Provide(controllers.NewGetClientsByCreatorController),
		fx.Provide(usecases.NewGetClientsByCreator),

		fx.Invoke(configureModuleRoutes),
	}
}
