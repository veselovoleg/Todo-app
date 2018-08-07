package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"go-echo-vue/models"

	"github.com/labstack/echo"
)

//H inerface
type H map[string]interface{}

//GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//Get tasks from db
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

//PutTask endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Create new task
		var task models.Task
		// Bind JSON to new task
		c.Bind(&task)
		// Add new task with model
		id, err := models.PutTask(db, task.Name)
		// if OK return JSON
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}

//DeleteTask endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use model for task delete
		_, err := models.DeleteTask(db, id)
		// if OK return JSON
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
