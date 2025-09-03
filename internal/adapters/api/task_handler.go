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
	r.GET("/tasks", auth, h.FetchTasks)
	r.PUT("/tasks/:id", auth, h.UpdateTask)
	r.DELETE("/tasks/:id", auth, h.DeleteTask)
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

func (h *TaskHandler) FetchTasks(c *gin.Context) {
	status := c.Query("status")
	userID := c.GetString("user_id")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tasks, err := h.service.FetchTasks(userID, status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id")

	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Status      string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	task, err := h.service.UpdateTask(id, userID, req.Title, req.Description, req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := h.service.DeleteTask(id, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}