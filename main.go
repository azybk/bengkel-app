package main

import (
	"bengkel_app/internal/component"
	"bengkel_app/internal/config"
	"bengkel_app/internal/module/customer"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.Get()
	dbConnection := component.GetDatabaseConnection(conf)

	customerRepository := customer.NewRepository(dbConnection)
	customerService := customer.NewService(customerRepository)

	app := fiber.New()
	customer.NewApi(app, customerService)

	_ = app.Listen(conf.Srv.Host + ":" + conf.Srv.Port)
}
