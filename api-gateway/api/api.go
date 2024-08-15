package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"github.com/dilshodforever/5-oyimtixon/api/handler"
	"github.com/dilshodforever/5-oyimtixon/api/middleware"
	_ "github.com/dilshodforever/5-oyimtixon/docs"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Gateway
// @version 1.0
// @description Dilshod's API Gateway
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// Middleware setup if needed
	ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}

	err = ca.LoadPolicy()
	if err != nil {
		panic(err)
	}
	router := r.Group("/")
	router.Use(middleware.NewAuth(ca))
	// Swagger documentation
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))

	// Account endpoints
	acc := router.Group("/account")
	{
		acc.POST("/create", h.CreateAccount)
		acc.GET("/get/:id", h.GetAccountById)
		acc.PUT("/update/:id", h.UpdateAccount)
		acc.DELETE("/delete/:id", h.DeleteAccount)
		acc.GET("/list", h.ListAccounts)
	}

	// Budget endpoints
	bud := router.Group("/budget")
	{
		bud.POST("/create", h.CreateBudget)
		bud.GET("/get/:id", h.GetBudgetByid)
		bud.PUT("/update/:id", h.UpdateBudget)
		bud.DELETE("/delete/:id", h.DeleteBudget)
		bud.GET("/list", h.ListBudgets)
	}

	// Category endpoints
	cat := router.Group("/category")
	{
		cat.POST("/create", h.CreateCategory)
		cat.PUT("/update/:id", h.UpdateCategory)
		cat.DELETE("/delete/:id", h.DeleteCategory)
		cat.GET("/list", h.ListCategories)
		cat.GET("/get/:id", h.GetByidCategory)
	}

	// Goal endpoints
	goa := router.Group("/goal")
	{
		goa.POST("/create", h.CreateGoal)
		goa.GET("/get/:id", h.GetGoalByid)
		goa.PUT("/update/:id", h.UpdateGoal)
		goa.DELETE("/delete/:id", h.DeleteGoal)
		goa.GET("/list", h.ListGoals)
	}

	// Transaction endpoints
	trans := router.Group("/transaction")
	{
		trans.POST("/create", h.CreateTransaction)
		trans.GET("/get/:id", h.GetTransaction)
		trans.PUT("/update/:id", h.UpdateTransaction)
		trans.DELETE("/delete/:id", h.DeleteTransaction)
		trans.GET("/list", h.ListTransactions)
	}

	notif := router.Group("/notifications")
	{
		notif.GET("/:id", h.GetNotification)
		notif.DELETE("/:id", h.DeleteNotification)
		notif.GET("/", h.ListNotification)
	}

	return r
}
