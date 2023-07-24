package transport

import (
	"strconv"

	"github.com/fiber-html/config"
	"github.com/fiber-html/internal/merchant/domain"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Merchant struct {
	Logger  config.Logger
	Usecase domain.MerchantExecutor
}

type MerchantHandler interface {
	GetAllMerchants(c *fiber.Ctx) error
	GetMerchantByID(c *fiber.Ctx) error
}

func NewMerchantHandler(usecase domain.MerchantExecutor, logger config.Logger) MerchantHandler {
	return &Merchant{
		Logger:  logger,
		Usecase: usecase,
	}
}

func (h *Merchant) GetAllMerchants(c *fiber.Ctx) error {
	var req domain.MerchantListRequest
	if err := c.QueryParser(&req); err != nil {
		h.Logger.Error("failed to scan request", zap.Error(err))
		return c.Render("index", fiber.Map{
			"Title":    "ERROR",
			"Subtitle": err.Error(),
		})
	}

	merchants, err := h.Usecase.GetMerchants(req)
	if err != nil {
		return c.Render("index", fiber.Map{
			"Title":    "ERROR",
			"Subtitle": err.Error(),
		})
	}

	return c.Render("views/merchant/_index", fiber.Map{
		"Title":     "Ini Data",
		"merchants": merchants,
	})
}

func (h *Merchant) GetMerchantByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		h.Logger.Error("failed to scan request", zap.Error(err))
		return c.Render("index", fiber.Map{
			"Title":    "ERROR",
			"Subtitle": err.Error(),
		})
	}

	merchant, err := h.Usecase.GetMerchantByID(id)
	if err != nil {
		return c.Render("index", fiber.Map{
			"Title":    "ERROR",
			"Subtitle": err.Error(),
		})
	}
	if merchant == nil {
		return c.Render("404", fiber.Map{
			"Title":    "ERROR",
			"Subtitle": err.Error(),
		})
	}

	return c.Render("views/merchant/_detail", fiber.Map{
		"Title":    "Ini Data",
		"merchant": merchant,
	})
}
