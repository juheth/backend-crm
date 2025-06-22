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
			{},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleProducts)
}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(dao.NewMySQLProductDao),
		fx.Provide(controllers.NewCreateProductController),
		fx.Provide(usecases.NewCreateProduct),

		fx.Invoke(configureModuleRoutes),
	}
}
