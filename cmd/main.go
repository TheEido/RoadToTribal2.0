package main

import (
	"RoadToTribal2.0/config"
	"RoadToTribal2.0/internal/adaptors/api"
	"RoadToTribal2.0/internal/adaptors/db"
	"RoadToTribal2.0/internal/repositories"
	"RoadToTribal2.0/internal/services/Transaction"
	"github.com/go-playground/validator/v10"
)

func main() {
	/*logger := config2.NewLogger()
	config := config2.LoadConfig(logger)
	logger.Infof("%v", config.Database)
	r := api.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))*/

	// Logger
	logger := config.NewLogger()
	defer config.CloseLogger(logger)

	// Configs
	configs := config.LoadConfig(logger)

	// Database Connection
	database := db.NewDatabaseConnection(logger, configs.Database)

	httpServer := api.NewHTTPServer(logger, configs.Server)

	// Validator
	validate := validator.New()

	transactionRepo := repositories.NewDatabaseRepository(logger, database)

	transactionSvc := Transaction.NewDefaultTransactionService(logger, transactionRepo)

	api.NewTransactionController(httpServer, logger, validate, transactionSvc)

	// -- End dependency injection section --

	// Let the party started!
	httpServer.Start()

}
