package main

import (
	"fmt"
	"log"

	"github.com/dilshodforever/5-oyimtixon/api"
	"github.com/dilshodforever/5-oyimtixon/api/handler"
	pba "github.com/dilshodforever/5-oyimtixon/genprotos/accaunts"
	pbb "github.com/dilshodforever/5-oyimtixon/genprotos/budgets"
	pbc "github.com/dilshodforever/5-oyimtixon/genprotos/categories"
	pbg "github.com/dilshodforever/5-oyimtixon/genprotos/goals"
	pbt "github.com/dilshodforever/5-oyimtixon/genprotos/transactions"
	pbn "github.com/dilshodforever/5-oyimtixon/genprotos/notifications"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	BudgetConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8087"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer BudgetConn.Close()
	Notifications, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8089"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer Notifications.Close()
	// Create service clients
	account := pba.NewAccountServiceClient(BudgetConn)
	budget := pbb.NewBudgetServiceClient(BudgetConn)
	category := pbc.NewCategoryServiceClient(BudgetConn)
	goal := pbg.NewGoalServiceClient(BudgetConn)
	transaction := pbt.NewTransactionServiceClient(BudgetConn)
	notifications:=pbn.NewNotificationtServiceClient(Notifications)
	// Create a new handler with the service clients
	h := handler.NewHandler(account, budget, category, goal, transaction, notifications)
	r := api.NewGin(h)
	fmt.Println("Server started on port:8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
