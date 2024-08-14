package handler

import (
	pba "github.com/dilshodforever/5-oyimtixon/genprotos/accaunts"
	pbb "github.com/dilshodforever/5-oyimtixon/genprotos/budgets"
	pbc "github.com/dilshodforever/5-oyimtixon/genprotos/categories"
    pbg "github.com/dilshodforever/5-oyimtixon/genprotos/goals"
    pbt "github.com/dilshodforever/5-oyimtixon/genprotos/transactions"
)

type Handler struct {
    Account   pba.AccountServiceClient
    Budget    pbb.BudgetServiceClient
    Category  pbc.CategoryServiceClient
    Goal      pbg.GoalServiceClient
    Transaction pbt.TransactionServiceClient
}

func NewHandler(
    account pba.AccountServiceClient, 
    budget pbb.BudgetServiceClient, 
    category pbc.CategoryServiceClient, 
    goal pbg.GoalServiceClient,
    transaction pbt.TransactionServiceClient,
) *Handler {
	return &Handler{
		Account:   account,
        Budget:    budget,
        Category:  category,
        Goal:      goal,
        Transaction: transaction,
	}
}
