// @title           Tech Challenge SNET API
// @version         1.0
// @description     API for establishments and stores management
// @host      localhost:8080
// @BasePath  /

package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/yMaatheus/tech-challenge-snet/config"
	_ "github.com/yMaatheus/tech-challenge-snet/docs"
	"github.com/yMaatheus/tech-challenge-snet/handler"
	"github.com/yMaatheus/tech-challenge-snet/repository"
	"github.com/yMaatheus/tech-challenge-snet/service"
	"go.uber.org/zap"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Setup Zap logger (production config)
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize zap logger: " + err.Error())
	}
	defer logger.Sync()

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		logger.Fatal("Failed to connect to the database", zap.Error(err))
	}
	defer db.Close()

	// Create Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Repository, Service and Handler initialization for Establishment
	establishmentRepo := repository.NewEstablishmentRepository(db)
	establishmentService := service.NewEstablishmentService(establishmentRepo)
	handler.NewEstablishmentHandler(e, establishmentService, logger)

	// Repository, Service and Handler initialization for Store
	storeRepo := repository.NewStoreRepository(db)
	storeService := service.NewStoreService(storeRepo)
	handler.NewStoreHandler(e, storeService, logger)

	// Swagger endpoint
	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.GET("/docs", func(c echo.Context) error {
		return c.Redirect(302, "/docs/index.html")
	})

	// Healthcheck (opcional)
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Info("Server running", zap.String("url", "http://localhost:"+port))
	e.Logger.Fatal(e.Start(":" + port))
}
