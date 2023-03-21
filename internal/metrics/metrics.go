package metrics

import (
	"database/sql"
	"log"
	database "operant/mysql"
)

type Metric struct {
	Name      string
	HighTreshold   int
	LowTreshold int
	Current    int
	ServiceId string
}

/*
SELECT A.`first_name` , A.`last_name` , B.`title`
FROM `members`AS A
INNER JOIN `movies` AS B
ON B.`id` = A.`movie_id`

SELECT M.*
FROM `Metrics`AS M
INNER JOIN `Services` AS S
ON M.`user_id` = S.`user_id`
*/

func GetMetricsByUser(userID string) []Metric {
	//stmt, err := database.Db.Prepare("select user_id from Services where user_id=? ;")
	stmt, err := database.Db.Prepare("SELECT M.* FROM `Metrics`AS M INNER JOIN `Services` AS S ON M.`service_id` = S.`service_id` where S.`user_id`=? ;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	metrics, err := ReadRows(rows)
	if err != nil{
		log.Fatal(err)
	}
	return metrics
}

func GetMetricsByUserFilterByTreshold(userID string, highTreshold int, lowTreshold int) []Metric {
	stmt, err := database.Db.Prepare("SELECT M.* FROM `Metrics`AS M INNER JOIN `Services` AS S ON M.`service_id` = S.`service_id` where S.`user_id`=? AND M.`high_treshold` <= ? AND M.`low_treshold` >= ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userID, highTreshold, lowTreshold)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	metrics, err := ReadRows(rows)
	if err != nil{
		log.Fatal(err)
	}
	return metrics
}

func ReadRows(rows *sql.Rows) ([]Metric, error) {
	var metrics []Metric
	for rows.Next() {
		var metric Metric
		err := rows.Scan(&metric.Name, &metric.HighTreshold, &metric.LowTreshold, &metric.Current, &metric.ServiceId)
		if err != nil{
			return nil, err
		}
		metrics = append(metrics, metric)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return metrics, nil;
}