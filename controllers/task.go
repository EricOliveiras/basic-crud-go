package controllers

import (
	"github/ericoliveiras/basic-crud-go/database"
	"github/ericoliveiras/basic-crud-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateTaskInput struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Finished    bool   `json:"finished"`
}

type UpdateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Finished    bool   `json:"finished"`
}

func CreateTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	input.ID = uuid.New().String()

	task := models.Task{
		ID:          input.ID,
		UserID:      input.UserID,
		Title:       input.Title,
		Description: input.Description,
		Finished:    false,
	}

	database.DB.Create(&task)

	c.JSON(http.StatusCreated, gin.H{"data": task})
}

func ReadTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func ReadTask(c *gin.Context) {
	var task models.Task
	if err := database.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ERROR::": "Task not found or not exists."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := database.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ERROR::": "Task not found or not exists."})
		return
	}

	var inputUpdateTask UpdateTaskInput
	if err := c.ShouldBindJSON(&inputUpdateTask); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	updateTask := models.Task {
		Title: inputUpdateTask.Title,
		Description: inputUpdateTask.Description,
		Finished: inputUpdateTask.Finished,
	}

	database.DB.Model(&task).Updates(&updateTask)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	if err := database.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ERROR::": "Task not found or not exists."})
		return
	}

	database.DB.Delete(&task)

	c.Status(http.StatusOK)
}