package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/fiber-html/config"
	"github.com/fiber-html/internal/merchant"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger := config.SetupLogger()
	loadEnv(logger)
	db, err := config.NewDBPool(
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_SCHEMA"),
		logger,
	)
	if err != nil {
		panic(err)
	}

	logger.Info("DB initiated", zap.Any("", db))

	engine := html.New("templates", ".gohtml")
	engine.AddFunc(
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views:         engine,
		Prefork:       true,
		CaseSensitive: true,
		ViewsLayout:   "layouts/main",
	})

	app.Static("/css", "templates/assets/css")
	app.Static("/img", "templates/assets/img")
	app.Static("/js", "templates/assets/js")
	app.Static("/vendor", "templates/assets/vendor")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":    "Test",
			"Subtitle": "kekekeke",
		})
	})

	merchant.Provide(app, db, logger)

	if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Fatal(err)
	}
}

func loadEnv(logger config.Logger) {
	_, err := os.Stat(".env")
	if err == nil {
		err = godotenv.Load()
		if err != nil {
			logger.Error("no .env files provided")
		}
	}
}
