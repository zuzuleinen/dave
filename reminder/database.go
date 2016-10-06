package reminder

import "database/sql"

type Reminder struct {
	Name string
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `CREATE TABLE IF NOT EXISTS reminders(Name TEXT);`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func Save(db *sql.DB, reminders []Reminder) {
	sql_additem := `INSERT INTO reminders(Name) values(?)`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range reminders {
		_, err2 := stmt.Exec(item.Name)
		if err2 != nil {
			panic(err2)
		}
	}
}


func Read(db *sql.DB) []Reminder {
	sql_readall := `SELECT Name FROM reminders`

	rows, err := db.Query(sql_readall)
	if err != nil { panic(err) }
	defer rows.Close()

	var result []Reminder
	for rows.Next() {
		item := Reminder{}
		err2 := rows.Scan(&item.Name)
		if err2 != nil { panic(err2) }
		result = append(result, item)
	}
	return result
}