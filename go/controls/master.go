package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"workout-note/models"
	"workout-note/services"

	"github.com/gin-gonic/gin"
)

func FetchMasterList(c *gin.Context) {
	user_param := c.Param("id")
	target_user, err := strconv.Atoi(user_param)
	if err != nil {
		target_user = 0
	}
	cache_key := "master_" + strconv.Itoa(target_user) + "_data"
	masters, cache_err := models.GetCache(cache_key)
	if cache_err != nil {
		fmt.Println("Access DB")
		masters, err = services.FetchMasterList(target_user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 2, "data": err.Error()})
			return
		}
		models.SetCache(cache_key, masters)
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": masters})
}

func CreateMaster(c *gin.Context) {
	var input models.Master
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	master, _ := services.CreateMaster(input)
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": master})
}

func UpdateMaster(c *gin.Context) {
	master_param := c.Param("id")
	target_master, err := strconv.Atoi(master_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 2, "data": err.Error()})
		return
	}
	var input models.Master
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 2, "data": err.Error()})
		return
	}
	master, _ := services.UpdateMaster(target_master, input)
	cache_key := "master_" + strconv.Itoa(master.UserID) + "_data"
	models.DeleteCache(cache_key)
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": master})
}

func DeleteMasterById(c *gin.Context) {
	master_param := c.Param("id")
	target_master, err := strconv.Atoi(master_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 2, "data": err.Error()})
		return
	}
	var master models.Master
	result, _ := services.DeleteMaster(target_master)
	cache_key := "master_" + strconv.Itoa(master.UserID) + "_data"
	models.DeleteCache(cache_key)
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": result})
}
