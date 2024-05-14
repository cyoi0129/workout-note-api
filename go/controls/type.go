package controllers

import (
	"fmt"
	"net/http"
	"workout-note/models"
	"workout-note/services"

	"github.com/gin-gonic/gin"
)

func FetchTypeList(c *gin.Context) {
	cache_key := "type_data"
	types, err := models.GetCache(cache_key)
	if err != nil {
		fmt.Println("Access DB")
		types, err = services.FetchTypeList()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 2, "data": err.Error()})
			return
		}
		models.SetCache(cache_key, types)
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": types})
}
