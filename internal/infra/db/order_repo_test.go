package db

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestMain(m *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		m.Error("erro de conexao com o banco")
	}

	db.Exec("CREATE TABLE orders (id VARCHAR(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL)")
}
