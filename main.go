package main

import (
	server "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/server"
	client "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente"
	Product "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto"
	user "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User"
)

func main() {
	app := server.ProvidersStore{}
	app.Init()
	app.AddModule(user.ModuleProviders())
	app.AddModule(client.ModuleProviders())
	app.AddModule(Product.ModuleProviders())
	app.Up()
}
