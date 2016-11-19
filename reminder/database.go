package reminder

import (
	"database/sql"
	"fmt"
)

type Reminder struct {
	Name      string
	Time      string
	Processed int
	RowId     int
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `CREATE TABLE IF NOT EXISTS reminders(Name TEXT, Time TEXT, Processed SMALLINT);`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

//Save all reminders in the db. Ignores the RowId field from Reminder
func Save(db *sql.DB, reminders []Reminder) {
	sql_additem := `INSERT INTO reminders(Name,Time,Processed) values(?,?,?)`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range reminders {
		_, err2 := stmt.Exec(item.Name, item.Time, item.Processed)
		if err2 != nil {
			panic(err2)
		}
	}
}

func Read(db *sql.DB) []Reminder {
	sql_readAll := `SELECT Name, Time, Processed, rowid FROM reminders`

	rows, err := db.Query(sql_readAll)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Reminder
	for rows.Next() {
		item := Reminder{}
		err2 := rows.Scan(&item.Name, &item.Time, &item.Processed, &item.RowId)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}

func Delete(rowId int, db *sql.DB) {
	sql_delete := fmt.Sprintf("delete from reminders where rowid = %v;", rowId)

	_, err := db.Exec(sql_delete)

	if err != nil {
		panic(err)
	}
}
