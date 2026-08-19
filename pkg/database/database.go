package database

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
)

var DB *goqu.Database
var pqconn *sql.DB

func Init(uri string) error {
	strconn := fmt.Sprintf("%v?sslmode=disable", uri)

	pqconn, err := sql.Open("postgres", strconn)
	if err != nil {
		return err
	}

	err = pqconn.Ping()
	if err != nil {
		return err
	}

	dialect := goqu.Dialect("postgres")
	DB = dialect.DB(pqconn)

	return nil
}

func Close() error {
	if pqconn == nil {
		return fmt.Errorf("Database already disconnect")
	}

	return pqconn.Close()
}
