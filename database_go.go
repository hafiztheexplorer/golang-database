package database_go

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root2adminthistimearound@tcp(localhost:3306)/belajar_golang_database") // return 2 value db dan err
	if err != nil {
		panic(err) // kalo error display errornya apa
	}

	// ini database poolingnya, ada 4 settings ini aja
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(90 * time.Minute)

	// menggunakan DB
	return db
}
