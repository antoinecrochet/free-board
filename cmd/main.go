package main

import (
	"log/slog"
	"os"

	"github.com/antoinecrochet/free-board/internal/adapter/primary/api"
	"github.com/antoinecrochet/free-board/internal/adapter/secondary/mariadb"
	"github.com/antoinecrochet/free-board/internal/core/service"
	"github.com/joho/godotenv"
)

func main() {
	slog.Info("Starting application...")
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	slog.Debug("DB Config", "user", dbUser, "password", dbPassword, "name", dbName)

	dataPort := mariadb.NewMariaDbProvider(dbUser, dbPassword, dbName)
	boardService := service.NewBoard(dataPort)

	app := api.NewApplication(boardService)
	app.StartServer()
}
