package main

import (
	"bengkel_app/internal/component"
	"bengkel_app/internal/config"
	"bengkel_app/internal/module/customer"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	conf := config.Get()
	dbConnection := component.GetDatabaseConnection(conf)

	customerRepository := customer.NewRepository(dbConnection)
	customerService := customer.NewService(customerRepository)

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${locals:requestid}] ${ip} - ${method} ${status} ${path}",
	}))
	customer.NewApi(app, customerService)

	_ = app.Listen(conf.Srv.Host + ":" + conf.Srv.Port)
}
