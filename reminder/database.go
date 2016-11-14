package reminder

import "database/sql"

type Reminder struct {
	Name string
	Time string
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `CREATE TABLE IF NOT EXISTS reminders(Name TEXT, Time TEXT);`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func Save(db *sql.DB, reminders []Reminder) {
	sql_additem := `INSERT INTO reminders(Name,Time) values(?,?)`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range reminders {
		_, err2 := stmt.Exec(item.Name, item.Time)
		if err2 != nil {
			panic(err2)
		}
	}
}

func Read(db *sql.DB) []Reminder {
	sql_readAll := `SELECT Name, Time FROM reminders`

	rows, err := db.Query(sql_readAll)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Reminder
	for rows.Next() {
		item := Reminder{}
		err2 := rows.Scan(&item.Name, &item.Time)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}
