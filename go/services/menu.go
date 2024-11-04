package services

import (
	"fmt"
	"workout-note-api/models"

	"github.com/lib/pq"
)

func FetchMenuList() ([]models.Menu, error) {
	var menus []models.Menu
	rows, err := models.DB.Query("SELECT id, name, image, type, target, muscles FROM \"workout_menus\"")
	if err != nil {
		fmt.Println(err)
	}
	var musclesStr []string
	for rows.Next() {
		var menu models.Menu
		rows.Scan(&menu.Id, &menu.Name, &menu.Image, &menu.Type, &menu.Target, pq.Array(&musclesStr))
		menu.Muscles = convert2Int(musclesStr)
		menus = append(menus, menu)
	}
	return menus, nil
}
