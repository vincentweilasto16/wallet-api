package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/vincentweilasto16/wallet-api/internal/config"
	"github.com/vincentweilasto16/wallet-api/internal/controller"
	"github.com/vincentweilasto16/wallet-api/internal/migration"
	repo "github.com/vincentweilasto16/wallet-api/internal/repository/postgres"
	"github.com/vincentweilasto16/wallet-api/internal/router"
	"github.com/vincentweilasto16/wallet-api/internal/service"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load DB config from env
	dbConfig := config.LoadDBConfig()

	// Connect to DB
	db, err := config.ConnectDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Run database migrations
	migration.RunDatabaseMigrations(db)

	// Initialize DB Repo
	dbRepo := repo.New(db)

	// Initialize service
	userService := service.NewUserService(dbRepo)
	transactionService := service.NewTransactionService(dbRepo)

	// Initialize controllers
	userController := controller.NewUserController(userService)
	transactionController := controller.NewTransactionController(transactionService)

	ctrl := &controller.Controllers{
		UserController:        userController,
		TransactionController: transactionController,
	}

	// Setup router
	r := router.NewRouter(ctrl)

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
