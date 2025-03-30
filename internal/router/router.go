package router

import (
	"fs.seaotterms.com-backend/internal/api"

	"github.com/gofiber/fiber/v2"
)

func ApiRouter(routerGroup fiber.Router) {
	routerGroup.Get("/directory", func(c *fiber.Ctx) error {
		return api.GetDirectory(c)
	})
}
