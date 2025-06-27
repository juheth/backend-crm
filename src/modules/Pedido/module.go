package Pedido

import (
	"net/http"

	types "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/types"
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/controllers"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/usecases"
	"go.uber.org/fx"
)

func configureModuleRoutes(
	ctrlCreateOrder *controllers.CreateOrderController,
	h *types.HandlersStore,
) {
	handlersModuleOrders := &types.SliceHandlers{
		Prefix: "orders",
		Routes: []types.HandlerModule{
			{
				Route:        "create",
				Method:       http.MethodPost,
				Handler:      ctrlCreateOrder.Run,
				RequiresAuth: true,
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleOrders)
}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(dao.NewMySQLOrderDao),
		fx.Provide(usecases.NewCreateOrder),
		fx.Provide(controllers.NewCreateOrderController),
		fx.Invoke(configureModuleRoutes),
	}
}
