package producto

import (
	"net/http"

	types "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/types"
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/controllers"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/usecases"

	"go.uber.org/fx"
)

func configureModuleRoutes(
	ctrlCreateProduct *controllers.CreateProductController,
	ctrlGetAllProducts *controllers.GetAllProductsController,
	ctrlGetProductByID *controllers.GetProductByIDController,
	ctrlUpdateProduct *controllers.UpdateProductController,
	ctrlDeactivateProduct *controllers.DeactivateProductController,

	h *types.HandlersStore,
) {
	handlersModuleProducts := &types.SliceHandlers{
		Prefix: "products",
		Routes: []types.HandlerModule{
			{
				Route:        "/create",
				Method:       http.MethodPost,
				Handler:      ctrlCreateProduct.Run,
				RequiresAuth: true,
			},
			{
				Route:        "/",
				Method:       http.MethodGet,
				Handler:      ctrlGetAllProducts.Run,
				RequiresAuth: true,
			},
			{
				Route:        "/:id",
				Method:       http.MethodGet,
				Handler:      ctrlGetProductByID.Run,
				RequiresAuth: true,
			},
			{
				Route:        "/Update/:id",
				Method:       http.MethodPut,
				Handler:      ctrlUpdateProduct.Run,
				RequiresAuth: true,
			},
			{
				Route:        "/:id/deactivate",
				Method:       http.MethodPut,
				Handler:      ctrlDeactivateProduct.Run,
				RequiresAuth: true,
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleProducts)
}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(dao.NewMySQLProductDao),
		fx.Provide(controllers.NewCreateProductController),
		fx.Provide(usecases.NewCreateProduct),
		fx.Provide(controllers.NewGetAllProductsController),
		fx.Provide(usecases.NewGetAllProducts),
		fx.Provide(controllers.NewGetProductByIDController),
		fx.Provide(usecases.NewGetProductByID),
		fx.Provide(controllers.NewUpdateProductController),
		fx.Provide(usecases.NewUpdateProduct),
		fx.Provide(controllers.NewDeactivateProductController),
		fx.Provide(usecases.NewDeactivateProduct),

		fx.Invoke(configureModuleRoutes),
	}
}
