package merchant

import (
	"github.com/fiber-html/config"
	"github.com/fiber-html/internal/merchant/domain"
	"github.com/fiber-html/internal/merchant/transport"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Provide(tp *fiber.App, db *pgxpool.Pool, logger config.Logger) {
	repository := domain.NewRepository(db, logger)
	useCase := domain.NewUseCase(repository, logger)
	transport.Route(tp, useCase, logger)
}
