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
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	keycloakServerUrl := os.Getenv("KEYCLOAK_SERVER_URL")
	keycloakRealm := os.Getenv("KEYCLOAK_REALM")

	if dbUser == "" || dbPassword == "" || dbName == "" || keycloakServerUrl == "" || keycloakRealm == "" {
		slog.Error("Missing required environment variables")
		return
	}

	slog.Debug("DB Config", "user", dbUser, "password", dbPassword, "name", dbName)

	dataPort := mariadb.NewMariaDbProvider(dbUser, dbPassword, dbName)
	boardService := service.NewBoard(dataPort)

	app := api.NewApplication(boardService)
	app.StartServer(keycloakServerUrl, keycloakRealm)
}
