package controllers

import (
	"fmt"
	"net/http"
	"workout-note-api/models"
	"workout-note-api/services"

	"github.com/gin-gonic/gin"
)

func FetchMasterList(c *gin.Context) {
	data, err := models.GetCache("master") // キャッシュに保存済みのデータを取得

	if err != nil { // キャッシュが存在しない場合、DBアクセスへ
		fmt.Println("Access DB")
		areas, err := services.FetchAreaList()
		lines, err := services.FetchLineList()
		gyms, err := services.FetchGymList()
		stations, err := services.FetchStationList()
		menus, err := services.FetchMenuList()
		muscles, err := services.FetchMuscleList()

		var line_stations []models.LineStation
		for _, line := range lines {
			target_stations, _ := services.FetchStationListByLine(line.Id)
			line_station := models.LineStation{
				Id:       line.Id,
				Name:     line.Name,
				Stations: target_stations,
			}
			line_stations = append(line_stations, line_station)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 2, "error": err.Error()})
			return
		}
		data = models.Master{
			Lines:    line_stations,
			Stations: stations,
			Gyms:     gyms,
			Areas:    areas,
			Menus:    menus,
			Muscles:  muscles,
		}

		models.SetCache("master", data) // DBから取得したデータをキャッシュへ保存
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": data})
}
