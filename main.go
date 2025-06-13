package main

import (
	user "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User"
	server "github.com/juheth/Go-Clean-Arquitecture/src/infrastructure/server"
)

func main() {
	app := server.ProvidersStore{}
	app.Init()
	app.AddModule(user.ModuleProviders())
	app.Up()
}
