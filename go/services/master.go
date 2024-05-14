package services

import (
	"fmt"
	"strconv"
	"workout-note/models"

	"github.com/lib/pq"
)

func convert2Int(strings []string) (ints []int) {
	var results []int
	for _, i := range strings {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		results = append(results, j)
	}
	return results
}

func FetchMasterList(user_id int) ([]models.Master, error) {
	var masters []models.Master
	rows, err := models.DB.Query("SELECT id, userID, name, image, type, target, muscles FROM \"masters\" WHERE userID = $1", user_id)
	if err != nil {
		fmt.Println(err)
	}
	var musclesStr []string
	for rows.Next() {
		var master models.Master
		rows.Scan(&master.Id, &master.UserID, &master.Name, &master.Image, &master.Type, &master.Target, pq.Array(&musclesStr))
		master.Muscles = convert2Int(musclesStr)
		masters = append(masters, master)
	}
	fmt.Printf("%v", masters)
	return masters, nil
}

func CreateMaster(input models.Master) (models.Master, error) {
	master := models.Master{
		UserID:  input.UserID,
		Name:    input.Name,
		Image:   input.Image,
		Type:    input.Type,
		Target:  input.Target,
		Muscles: input.Muscles,
	}
	err := models.DB.QueryRow("INSERT INTO masters(userID, name, image, type, target, muscles) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", master.UserID, master.Name, master.Image, master.Type, master.Target, master.Muscles).Scan(&master.Id)
	if err != nil {
		fmt.Println(err)
		return master, err
	}
	return master, nil
}

func UpdateMaster(master_id int, input models.Master) (models.Master, error) {
	master := models.Master{
		Id:      input.Id,
		UserID:  input.UserID,
		Name:    input.Name,
		Image:   input.Image,
		Type:    input.Type,
		Target:  input.Target,
		Muscles: input.Muscles,
	}
	_, err := models.DB.Query("UPDATE \"masters\" SET userID = $1, name = $2, image = $3 type = $4 target = $5 muscles = $6 WHERE id = $7", master.UserID, master.Name, master.Image, master.Type, master.Target, master.Muscles, master_id)
	if err != nil {
		return master, err
	}
	return master, nil
}

func DeleteMaster(master_id int) (bool, error) {
	_, err := models.DB.Query("DELETE FROM \"masters\" WHERE id = $1", master_id)
	if err != nil {
		return false, err
	}
	return true, nil
}
