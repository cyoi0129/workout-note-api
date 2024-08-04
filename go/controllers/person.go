package controllers

import (
	"fmt"
	"net/http"

	"strconv"
	"workout-note-api/models"
	"workout-note-api/services"

	"github.com/gin-gonic/gin"
)

// 指定条件のパーソン取得
func FetchTargetPersons(c *gin.Context) {
	var filter models.PersonFilter
	var persons []models.Person
	var err error
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "data": err.Error()})
		return
	}
	fmt.Println(len(filter.Areas), len(filter.Stations))
	if len(filter.Areas) == 0 && len(filter.Stations) == 0 {
		persons, err = services.FetchTargetPersonsByGyms(filter)
	} else if len(filter.Stations) == 0 {
		persons, err = services.FetchTargetPersonsByAreas(filter)
	} else if len(filter.Areas) == 0 {
		persons, err = services.FetchTargetPersonsByStations(filter)
	} else {
		persons, err = services.FetchTargetPersons(filter)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": persons})
	}
}

// 既存パーソンの更新
func UpdatePerson(c *gin.Context) {
	user_param := c.Param("id")
	id, err := strconv.Atoi(user_param)
	if err != nil {
		id = 0
	}
	var input models.Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "data": err.Error()})
		return
	}

	person, db_error := services.UpdatePerson(id, input)
	if db_error != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": person})
	}
}

// パーソン情報の取得
func FetchPersonByID(c *gin.Context) {
	user_param := c.Param("id")
	id, _ := strconv.Atoi(user_param)
	person, err := services.FetchPersonByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": person})
	}
}
