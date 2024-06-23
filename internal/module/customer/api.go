package customer

import (
	"bengkel_app/domain"

	"github.com/gofiber/fiber/v2"
)

type api struct {
	customerService domain.CustomerService
}

func NewApi(app *fiber.App, customerService domain.CustomerService) {
	api := api{
		customerService,
	}

	app.Get("/foo/bar", api.FooBar)
}

func (a api) FooBar(ctx *fiber.Ctx) error {
	return ctx.JSON("OKE..")
}
