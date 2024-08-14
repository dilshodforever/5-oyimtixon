package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/dilshodforever/5-oyimtixon/genprotos/goals"
)

// CreateGoal handles creating a new goal
// @Summary      Create Goal
// @Description  Create a new goal
// @Tags         Goal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        goal body pb.CreateGoalRequest true "Goal details"
// @Success      200 {object} pb.GoalResponse "Goal created successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while creating goal"
// @Router       /goal/create [post]
func (h *Handler) CreateGoal(ctx *gin.Context) {
	var req pb.CreateGoalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	res, err := h.Goal.CreateGoal(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetGoalByid handles retrieving a goal by ID
// @Summary      Get Goal by ID
// @Description  Retrieve details of a goal by ID
// @Tags         Goal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id query string true "Goal ID"
// @Success      200 {object} pb.GetGoalResponse "Goal details"
// @Failure      400 {string} string "Missing or invalid ID"
// @Failure      500 {string} string "Error while fetching goal"
// @Router       /goal/get [get]
func (h *Handler) GetGoalByid(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "Missing or invalid ID"})
		return
	}

	req := &pb.GetGoalByidRequest{Id: id}

	res, err := h.Goal.GetGoalByid(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// UpdateGoal handles updating a goal
// @Summary      Update Goal
// @Description  Update details of a goal
// @Tags         Goal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        goal body pb.UpdateGoalRequest true "Updated goal details"
// @Success      200 {object} pb.GoalResponse "Goal updated successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while updating goal"
// @Router       /goal/update [put]
func (h *Handler) UpdateGoal(ctx *gin.Context) {
	var req pb.UpdateGoalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	res, err := h.Goal.UpdateGoal(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// DeleteGoal handles deleting a goal
// @Summary      Delete Goal
// @Description  Delete a goal by ID
// @Tags         Goal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id query string true "Goal ID"
// @Success      200 {object} pb.GoalResponse "Goal deleted successfully"
// @Failure      400 {string} string "Missing or invalid ID"
// @Failure      500 {string} string "Error while deleting goal"
// @Router       /goal/delete [delete]
func (h *Handler) DeleteGoal(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "Missing or invalid ID"})
		return
	}

	req := &pb.DeleteGoalRequest{Id: id}

	res, err := h.Goal.DeleteGoal(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// ListGoals handles listing all goals
// @Summary      List Goals
// @Description  Get a list of all goals
// @Tags         Goal
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.ListGoalsResponse "List of goals"
// @Failure      500 {string} string "Error while fetching goals"
// @Router       /goal/list [get]
func (h *Handler) ListGoals(ctx *gin.Context) {
	req := &pb.ListGoalsRequest{}

	res, err := h.Goal.ListGoals(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}
