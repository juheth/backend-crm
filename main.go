package main

import (
	server "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/server"
	client "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente"
	user "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User"
)

func main() {
	app := server.ProvidersStore{}
	app.Init()
	app.AddModule(user.ModuleProviders())
	app.AddModule(client.ModuleProviders())
	app.Up()
}
