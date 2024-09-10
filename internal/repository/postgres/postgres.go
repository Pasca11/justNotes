package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type Database struct {
	DB *sqlx.DB
}

func NewDatabase() (*Database, error) {
	dataBase := &Database{}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	dataBase.DB = db
	err = dataBase.CreateNewUserTable()
	return dataBase, err
}

func (db *Database) CreateNewUserTable() error {
	newTableString := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL
	);`

	_, err := db.DB.Exec(newTableString)
	return err
}
