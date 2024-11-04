package services

import (
	"fmt"
	"workout-note-api/models"
)

func FetchStationList() ([]models.Station, error) {
	var stations []models.Station
	rows, err := models.DB.Query("SELECT id, lineID, name FROM \"workout_stations\"")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var station models.Station
		rows.Scan(&station.Id, &station.LineID, &station.Name)
		stations = append(stations, station)
	}
	return stations, nil
}

func FetchStationListByLine(line_id uint) ([]models.Station, error) {
	var stations []models.Station
	rows, err := models.DB.Query("SELECT id, lineID, name FROM \"workout_stations\" WHERE lineID = $1", line_id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var station models.Station
		rows.Scan(&station.Id, &station.LineID, &station.Name)
		stations = append(stations, station)
	}
	return stations, nil
}

func FetchAreaList() ([]models.Area, error) {
	var areas []models.Area
	rows, err := models.DB.Query("SELECT id, name FROM \"workout_areas\"")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var area models.Area
		rows.Scan(&area.Id, &area.Name)
		areas = append(areas, area)
	}
	return areas, nil
}

func FetchGymList() ([]models.Gym, error) {
	var gyms []models.Gym
	rows, err := models.DB.Query("SELECT id, name FROM \"workout_gyms\"")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var gym models.Gym
		rows.Scan(&gym.Id, &gym.Name)
		gyms = append(gyms, gym)
	}
	return gyms, nil
}

func FetchLineList() ([]models.Line, error) {
	var lines []models.Line
	rows, err := models.DB.Query("SELECT id, name FROM \"workout_lines\"")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var line models.Line
		rows.Scan(&line.Id, &line.Name)
		lines = append(lines, line)
	}
	return lines, nil
}
