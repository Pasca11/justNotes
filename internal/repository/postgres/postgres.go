package postgres

import (
	"fmt"
	"github.com/Pasca11/justNotes/models"
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
	if err != nil {
		return nil, err
	}
	err = dataBase.CreateNewNoteTable()
	if err != nil {
		return nil, err
	}
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

func (db *Database) CreateNewNoteTable() error {
	newTableString := `
		CREATE TABLE IF NOT EXISTS notes (
 	    id SERIAL PRIMARY KEY,
 	    text VARCHAR NOT NULL,
 		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 		deadline TIMESTAMP,
 		user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE
	);`
	_, err := db.DB.Exec(newTableString)
	return err
}

func (db *Database) CreateUser(user *models.User) error {
	stmt := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err := db.DB.Exec(stmt, user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("repo.insert.user: %w", err)
	}
	return nil
}

func (db *Database) GetUser(username string) (*models.User, error) {
	user := &models.User{}
	err := db.DB.Get(user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		return nil, fmt.Errorf("repo.get.user: %w", err)
	}
	return user, nil
}

func (db *Database) GetNotes(userId int) ([]models.Note, error) {
	//user, err := db.GetUser(username)
	//if err != nil {
	//	return nil, fmt.Errorf("repo.get.notes: %w", err)
	//}
	var notes []models.Note
	err := db.DB.Select(&notes, "SELECT * FROM notes WHERE user_id=$1", userId)
	if err != nil {
		return nil, fmt.Errorf("repo.get.notes: %w", err)
	}
	return notes, nil
}

func (db *Database) CreateNote(id int, note *models.Note) error {
	const op = "repo.create.note"
	stmt := "INSERT INTO notes (text, user_id, deadline) VALUES ($1, $2, $3)"
	_, err := db.DB.Exec(stmt, note.Text, id, note.Deadline)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (db *Database) DeleteNote(noteId int) error {
	op := "repo.delete.note"
	stmt := "DELETE FROM notes WHERE id=$1"
	_, err := db.DB.Exec(stmt, noteId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
