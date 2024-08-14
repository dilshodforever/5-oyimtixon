package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/dilshodforever/5-oyimtixon/genprotos/accaunts"
)

// CreateAccount handles creating a new account
// @Summary      Create Account
// @Description  Create a new account
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        account body pb.CreateAccountRequest true "Account details"
// @Success      200 {object} pb.CreateAccountResponse "Account created successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while creating account"
// @Router       /account/create [post]
func (h *Handler) CreateAccount(ctx *gin.Context) {
	var req pb.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	res, err := h.Account.CreateAccount(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetAccountById handles retrieving an account by ID
// @Summary      Get Account by ID
// @Description  Get details of an account by ID
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id query string true "Account ID"
// @Success      200 {object} pb.GetAccountByidResponse "Account details"
// @Failure      400 {string} string "Missing or invalid ID"
// @Failure      500 {string} string "Error while fetching account"
// @Router       /account/get [get]
func (h *Handler) GetAccountById(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "Missing or invalid ID"})
		return
	}

	req := &pb.GetByIdAccauntRequest{Id: id}

	res, err := h.Account.GetAccountByid(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// UpdateAccount handles updating an account
// @Summary      Update Account
// @Description  Update details of an account
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        account body pb.UpdateAccountRequest true "Updated account details"
// @Success      200 {object} pb.UpdateAccountResponse "Account updated successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while updating account"
// @Router       /account/update [put]
func (h *Handler) UpdateAccount(ctx *gin.Context) {
	var req pb.UpdateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	res, err := h.Account.UpdateAccount(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// DeleteAccount handles deleting an account
// @Summary      Delete Account
// @Description  Delete an account by ID
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id query string true "Account ID"
// @Success      200 {object} pb.UpdateAccountResponse "Account deleted successfully"
// @Failure      400 {string} string "Missing or invalid ID"
// @Failure      500 {string} string "Error while deleting account"
// @Router       /account/delete [delete]
func (h *Handler) DeleteAccount(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "Missing or invalid ID"})
		return
	}

	req := &pb.DeleteAccountRequest{Id: id}

	res, err := h.Account.DeleteAccount(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// ListAccounts handles listing all accounts
// @Summary      List Accounts
// @Description  Get a list of all accounts
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.ListAccountsResponse "List of accounts"
// @Failure      500 {string} string "Error while fetching accounts"
// @Router       /account/list [get]
func (h *Handler) ListAccounts(ctx *gin.Context) {
	req := &pb.ListAccountsRequest{}

	res, err := h.Account.ListAccounts(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}
