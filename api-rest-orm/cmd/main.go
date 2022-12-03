package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	api_rest_orm "orm"
	"orm/pkg/logger"
	"orm/pkg/server"
	"orm/pkg/services"
	"orm/pkg/storage/postgresql"
	"os"
)

func main() {
	var port, logLevel, storage string
	var repo api_rest_orm.Repository
	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}
	if storage = os.Getenv("STORAGE"); storage == "" {
		storage = "postgresql"
	}

	logLevel = os.Getenv("LOG_LEVEL")
	log := logger.NewLogger(logLevel)

	switch storage {
	case "postgresql":
		db, _ := postgresql.ConnDB()
		postgresql.ConfigShema(&api_rest_orm.Producto{})
		repo = postgresql.NewRepository(db)
	}

	serviceProducto := services.NewService(repo)

	app := fiber.New(fiber.Config{AppName: "api-rest-orm"})

	uuid, _ := uuid.DefaultGenerator.NewV4()

	server.NewServer(uuid.String(), log, app, serviceProducto)

	app.Listen(fmt.Sprintf(":%v", port))

}
