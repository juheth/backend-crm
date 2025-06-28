package pedido

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
	ctrlGetAllOrders *controllers.GetAllOrdersController,
	ctrlGetOrderByID *controllers.GetOrderByIDController,
	ctrlUpdateOrderStatus *controllers.UpdateOrderStatusController,
	ctrlDeleteOrder *controllers.DeleteOrderController,
	ctrlGetOrdersByClient *controllers.GetOrdersByClientController,
	ctrlGetOrdersByStatus *controllers.GetOrdersByStatusController,
	h *types.HandlersStore,
) {
	handlersModuleOrders := &types.SliceHandlers{
		Prefix: "orders",
		Routes: []types.HandlerModule{
			{
				Route:        "get-all",
				Method:       http.MethodGet,
				Handler:      ctrlGetAllOrders.Run,
				RequiresAuth: true,
			},
			{
				Route:        "by-status",
				Method:       http.MethodGet,
				Handler:      ctrlGetOrdersByStatus.Run,
				RequiresAuth: true,
			},
			{
				Route:        ":id",
				Method:       http.MethodGet,
				Handler:      ctrlGetOrderByID.Run,
				RequiresAuth: true,
			},
			{
				Route:        "by-client/:clientId",
				Method:       http.MethodGet,
				Handler:      ctrlGetOrdersByClient.Run,
				RequiresAuth: true,
			},
			{
				Route:        "create",
				Method:       http.MethodPost,
				Handler:      ctrlCreateOrder.Run,
				RequiresAuth: true,
			},
			{
				Route:        ":id/status",
				Method:       http.MethodPut,
				Handler:      ctrlUpdateOrderStatus.Run,
				RequiresAuth: true,
			},
			{
				Route:        ":id",
				Method:       http.MethodDelete,
				Handler:      ctrlDeleteOrder.Run,
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
		fx.Provide(usecases.NewGetAllOrders),
		fx.Provide(controllers.NewGetAllOrdersController),
		fx.Provide(usecases.NewGetOrderByID),
		fx.Provide(controllers.NewGetOrderByIDController),
		fx.Provide(usecases.NewUpdateOrderStatus),
		fx.Provide(controllers.NewUpdateOrderStatusController),
		fx.Provide(usecases.NewDeleteOrder),
		fx.Provide(controllers.NewDeleteOrderController),
		fx.Provide(usecases.NewGetOrdersByClient),
		fx.Provide(controllers.NewGetOrdersByClientController),
		fx.Provide(usecases.NewGetOrdersByStatus),
		fx.Provide(controllers.NewGetOrdersByStatusController),
		fx.Invoke(configureModuleRoutes),
	}
}
