package cleaner

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var (
	id      int
	created string
)

type entry struct {
	id      int
	created string
}

//CleanUp starts the cleaning process of entries older then the given retention period.
func CleanUp(db *sql.DB, retention int, metrics bool) {

	entries := getIDs(db, retention)

	copyEntries(db, entries)

	removeEntries(db, metrics)
}

func getIDs(db *sql.DB, retention int) []int {

	dt := time.Now().AddDate(0, 0, -retention).Format("2006-01-02")

	log.Info("Searching for entries before: ", dt)

	rows, err := db.Query("SELECT id,created FROM api_review WHERE created < $1", dt)
	if err != nil {
		log.Error("Error executing query ", err)
	}
	defer rows.Close()

	var result []int
	for rows.Next() {
		err := rows.Scan(&id, &created)
		if err != nil {
			log.Error(err)
		}
		result = append(result, id)
	}
	return result
}

func copyEntries(db *sql.DB, ids []int) {

	_, err := db.Exec(fmt.Sprintf("CREATE TABLE %s(id INT PRIMARY KEY NOT NULL)", "rows_to_delete"))
	if err != nil {
		log.Error(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Error(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(pq.CopyIn("rows_to_delete", "id"))
	if err != nil {
		log.Error(err)
	}

	for _, id := range ids {
		_, err = stmt.Exec(id)
		if err != nil {
			log.Error(err)
		}
	}

	stmt.Exec()
	if err != nil {
		log.Error(err)
	}
	defer stmt.Close()

	err = tx.Commit()
	if err != nil {
		log.Error(err)
	}
}

func removeEntries(db *sql.DB, metrics bool) {

	var tables []string
	if metrics {
		tables = append(tables, "rule_violation", "custom_label_mapping", "api_review")
	} else {
		tables = append(tables, "rule_violation", "api_review")
	}

	for _, table := range tables {

		tableID := "id"
		if table != "api_review" {
			tableID = "api_review_id"
		}

		log.Info("Begin cleaning up: ", table)

		_, err := db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s IN (SELECT id FROM rows_to_delete)", table, tableID))
		if err != nil {
			log.Error(err)
		}
	}

	_, err := db.Exec("DROP TABLE rows_to_delete")
	if err != nil {
		log.Error(err)
	}

	log.Info("Done Removing")
}
