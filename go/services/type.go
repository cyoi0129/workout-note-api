package services

import (
	"fmt"
	"workout-note/models"
)

func FetchTypeList() ([]models.Type, error) {
	var types []models.Type
	rows, err := models.DB.Query("SELECT id, name FROM \"types\"")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var item models.Type
		rows.Scan(&item.Id, &item.Name)
		types = append(types, item)
	}
	fmt.Printf("%v", types)
	return types, nil
}
