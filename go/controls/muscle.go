package controllers

import (
	"fmt"
	"net/http"
	"workout-note/models"
	"workout-note/services"

	"github.com/gin-gonic/gin"
)

func FetchMuscleList(c *gin.Context) {
	cache_key := "muscle_data"
	muscles, err := models.GetCache(cache_key)
	if err != nil {
		fmt.Println("Access DB")
		muscles, err = services.FetchMuscleList()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 2, "data": err.Error()})
			return
		}
		models.SetCache(cache_key, muscles)
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": muscles})
}
