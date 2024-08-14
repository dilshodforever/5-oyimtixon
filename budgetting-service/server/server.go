package server

import (
	"log"
	"net"

	pb "github.com/dilshodforever/5-oyimtixon/genprotos/accaunts"
	goalsPb "github.com/dilshodforever/5-oyimtixon/genprotos/goals"
	transactionsPb "github.com/dilshodforever/5-oyimtixon/genprotos/transactions"
	budgetting "github.com/dilshodforever/5-oyimtixon/genprotos/budgets"
	category "github.com/dilshodforever/5-oyimtixon/genprotos/categories"
	"github.com/dilshodforever/5-oyimtixon/service"
	postgres "github.com/dilshodforever/5-oyimtixon/storage/mongo"
	"google.golang.org/grpc"
)

func Connection() net.Listener {
	// Initialize MongoDB connection
	db, err := postgres.NewMongoConnection()
	if err != nil {
		log.Fatal("Error while connecting to DB: ", err.Error())
	}

	// Set up TCP listener
	lis, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatal("Error while starting TCP listener: ", err.Error())
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Initialize services
	accountService := service.NewAccountService(db)
	goalService := service.NewGoalService(db)
	transactionService := service.NewTransactionService(db)
	budget:= service.NewBudgetService(db)
	categories:=service.NewCategoryService(db)
	// Register services with the gRPC server
	pb.RegisterAccountServiceServer(s, accountService)
	goalsPb.RegisterGoalServiceServer(s, goalService)
	transactionsPb.RegisterTransactionServiceServer(s, transactionService)
	budgetting.RegisterBudgetServiceServer(s, budget)
	category.RegisterCategoryServiceServer(s, categories)
	// Start serving
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	return lis
}
