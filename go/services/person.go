package services

import (
	"fmt"
	"workout-note-api/models"

	"github.com/lib/pq"
)

// パーソン情報の取得
func FetchPersonByID(user_id int) (models.Person, error) {
	var person models.Person
	row := models.DB.QueryRow("SELECT id, userID, name, gender, brith, stations, areas, gyms, times, bp, sq, dl FROM \"workout_persons\" WHERE userID = $1", user_id)
	var stationStr, areaStr, gymStr []string
	err := row.Scan(&person.Id, &person.UserID, &person.Name, &person.Gender, &person.Brith, pq.Array(&stationStr), pq.Array(&areaStr), pq.Array(&gymStr), pq.Array(&person.Times), &person.Bp, &person.Sq, &person.Dl)
	person.Stations = convert2Int(stationStr)
	person.Areas = convert2Int(areaStr)
	person.Gyms = convert2Int(gymStr)

	if err != nil {
		fmt.Println(err)
		return person, err
	}
	return person, nil
}

// パーソン情報の更新
func UpdatePerson(user_id int, input models.Person) (models.Person, error) {
	person := models.Person{
		Id:       input.Id,
		UserID:   input.UserID,
		Name:     input.Name,
		Gender:   input.Gender,
		Brith:    input.Brith,
		Stations: input.Stations,
		Areas:    input.Areas,
		Gyms:     input.Gyms,
		Times:    input.Times,
		Bp:       input.Bp,
		Sq:       input.Sq,
		Dl:       input.Dl,
	}
	_, err := models.DB.Query("UPDATE \"workout_persons\" SET name = $1, gender = $2, brith = $3, stations = ($4), areas = ($5), gyms = ($6), times = ($7), bp = $8, sq = $9, dl = $10 WHERE userId = $11", person.Name, person.Gender, person.Brith, pq.Array(person.Stations), pq.Array(person.Areas), pq.Array(person.Gyms), pq.Array(person.Times), person.Bp, person.Sq, person.Dl, user_id)
	if err != nil {
		fmt.Println(err)
		return person, err
	}
	return person, nil
}

// 指定条件のパーソン取得
func FetchTargetPersons(filter models.PersonFilter) ([]models.Person, error) {
	var persons []models.Person
	rows, err := models.DB.Query("SELECT id, userID, name, gender, brith, stations, areas, gyms, times, bp, sq, dl FROM \"workout_persons\" WHERE (gyms && ($1)) AND (stations && ($2)) AND (areas && ($3))", pq.Array(filter.Gyms), pq.Array(filter.Stations), pq.Array(filter.Areas))
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var person models.Person
		var stationStr, areaStr, gymStr []string
		rows.Scan(&person.Id, &person.UserID, &person.Name, &person.Gender, &person.Brith, pq.Array(&stationStr), pq.Array(&areaStr), pq.Array(&gymStr), pq.Array(&person.Times), &person.Bp, &person.Sq, &person.Dl)
		person.Stations = convert2Int(stationStr)
		person.Areas = convert2Int(areaStr)
		person.Gyms = convert2Int(gymStr)
		persons = append(persons, person)
	}
	return persons, nil
}

// ジム指定のパーソン取得
func FetchTargetPersonsByGyms(filter models.PersonFilter) ([]models.Person, error) {
	var persons []models.Person
	rows, err := models.DB.Query("SELECT id, userID, name, gender, brith, stations, areas, gyms, times, bp, sq, dl FROM \"workout_persons\" WHERE gyms && ($1))", pq.Array(filter.Gyms))

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var person models.Person
		var stationStr, areaStr, gymStr []string
		rows.Scan(&person.Id, &person.UserID, &person.Name, &person.Gender, &person.Brith, pq.Array(&stationStr), pq.Array(&areaStr), pq.Array(&gymStr), pq.Array(&person.Times), &person.Bp, &person.Sq, &person.Dl)
		person.Stations = convert2Int(stationStr)
		person.Areas = convert2Int(areaStr)
		person.Gyms = convert2Int(gymStr)
		persons = append(persons, person)
	}
	return persons, nil
}

// 駅指定のパーソン取得
func FetchTargetPersonsByStations(filter models.PersonFilter) ([]models.Person, error) {
	var persons []models.Person
	rows, err := models.DB.Query("SELECT id, userID, name, gender, brith, stations, areas, gyms, times, bp, sq, dl FROM \"workout_persons\" WHERE (gyms && ($1)) AND (stations && ($2))", pq.Array(filter.Gyms), pq.Array(filter.Stations))

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var person models.Person
		var stationStr, areaStr, gymStr []string
		rows.Scan(&person.Id, &person.UserID, &person.Name, &person.Gender, &person.Brith, pq.Array(&stationStr), pq.Array(&areaStr), pq.Array(&gymStr), pq.Array(&person.Times), &person.Bp, &person.Sq, &person.Dl)
		person.Stations = convert2Int(stationStr)
		person.Areas = convert2Int(areaStr)
		person.Gyms = convert2Int(gymStr)
		persons = append(persons, person)
	}
	return persons, nil
}

// エリア指定のパーソン取得
func FetchTargetPersonsByAreas(filter models.PersonFilter) ([]models.Person, error) {
	var persons []models.Person
	rows, err := models.DB.Query("SELECT id, userID, name, gender, brith, stations, areas, gyms, times, bp, sq, dl FROM \"workout_persons\" WHERE (gyms && ($1)) AND (areas && ($2))", pq.Array(filter.Gyms), pq.Array(filter.Areas))
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var person models.Person
		var stationStr, areaStr, gymStr []string
		rows.Scan(&person.Id, &person.UserID, &person.Name, &person.Gender, &person.Brith, pq.Array(&stationStr), pq.Array(&areaStr), pq.Array(&gymStr), pq.Array(&person.Times), &person.Bp, &person.Sq, &person.Dl)
		person.Stations = convert2Int(stationStr)
		person.Areas = convert2Int(areaStr)
		person.Gyms = convert2Int(gymStr)
		persons = append(persons, person)
	}
	return persons, nil
}
