package credentials

import "database/sql"

type Credential struct {
	 Username string
	 Password string
	 Website string
	 Notes string
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `CREATE TABLE IF NOT EXISTS credentials(Username TEXT, Password TEXT, Website TEXT, Notes TEXT);`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func Save(db *sql.DB, credentials []Credential) {
	sql_additem := `INSERT INTO credentials(Username, Password, Website, Notes) values(?,?,?,?)`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range credentials {
		_, err2 := stmt.Exec(item.Username, item.Password, item.Website, item.Notes)
		if err2 != nil {
			panic(err2)
		}
	}
}

func Read(db *sql.DB) []Credential {
	sql_readall := `SELECT Username, Password, Website, Notes FROM credentials`

	rows, err := db.Query(sql_readall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Credential
	for rows.Next() {
		item := Credential{}
		err2 := rows.Scan(&item.Username, &item.Password, &item.Website, &item.Notes)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}
