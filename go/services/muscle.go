package services

import (
	"fmt"
	"workout-note-api/models"
)

func FetchMuscleList() ([]models.Muscle, error) {
	var muscles []models.Muscle
	rows, err := models.DB.Query("SELECT id, part, name FROM \"workout_muscles\"")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var muscle models.Muscle
		rows.Scan(&muscle.Id, &muscle.Part, &muscle.Name)
		muscles = append(muscles, muscle)
	}
	return muscles, nil
}
