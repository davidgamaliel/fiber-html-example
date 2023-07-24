package transport

import (
	"github.com/fiber-html/config"
	"github.com/fiber-html/internal/merchant/domain"
	"github.com/gofiber/fiber/v2"
)

func Route(tp *fiber.App, uc domain.MerchantExecutor, logger config.Logger) {
	handler := NewMerchantHandler(uc, logger)
	merch := tp.Group("/merchants")

	merch.Get("/", handler.GetAllMerchants)
	merch.Get("/:id", handler.GetMerchantByID)
}
