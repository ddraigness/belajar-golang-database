package belajargolangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// func TestEmpty(t *testing.T) {
	
// }

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:P@ssw0rd@tcp(localhost:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}