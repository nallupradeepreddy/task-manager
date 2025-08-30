package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nallupradeepreddy/task-manager/internal/core/ports"
)

type TaskHandler struct {
	service ports.TaskService
}

func NewTaskHandler(service ports.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) RegisterRoutes(r *gin.Engine, auth gin.HandlerFunc) {
	r.POST("/tasks", auth, h.CreateTask)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	task, err := h.service.CreateTask(userID, req.Title, req.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}
