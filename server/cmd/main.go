// @title           Tech Challenge SNET API
// @version         1.0
// @description     API for establishments and stores management
// @host      localhost:8080
// @BasePath  /

package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/yMaatheus/tech-challenge-snet/config"
	_ "github.com/yMaatheus/tech-challenge-snet/docs"
	"github.com/yMaatheus/tech-challenge-snet/handler"
	"github.com/yMaatheus/tech-challenge-snet/repository"
	"github.com/yMaatheus/tech-challenge-snet/service"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Repository, Service and Handler initialization for Establishment
	establishmentRepo := repository.NewEstablishmentRepository(db)
	establishmentService := service.NewEstablishmentService(establishmentRepo)

	// Create Echo instance
	e := echo.New()

	// Register Establishment endpoints
	handler.NewEstablishmentHandler(e, establishmentService)

	// Swagger endpoint
	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.GET("/docs", func(c echo.Context) error {
		return c.Redirect(302, "/docs/index.html")
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
