package database_test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestKosongan(t *testing.T) {

}

func TestOpenConnectionDatabase(t *testing.T) {
	db, err := sql.Open("mysql", "root:root2adminthistimearound@tcp(localhost:3306)/belajar_golang_database") // return 2 value db dan err
	if err != nil {
		panic(err) // kalo error display errornya apa
	}

	defer db.Close() // menutup object sql.DB
	// menggunakan DB
}
