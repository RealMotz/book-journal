package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	DB *sql.DB
}

type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Memo        string `json:"memo"`
	IsReading   bool   `json:"isReading"`
}

func (sqlite *SqliteDB) InitializeDB() {
	database, err := sql.Open("sqlite3", "./books.db")

	if err != nil {
		log.Fatal(err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS books (
    id          INTEGER PRIMARY KEY,
    name        TEXT NOT NULL,
    description Text,
    memo        TEXT,
    isReading   BOOLEAN);
  `
	_, err = database.Exec(query)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
	}

	sqlite.DB = database
}

func (sqlite *SqliteDB) Insert(book Book) error {
	query := `
  INSERT INTO books (name, description, memo, isReading)
  VALUES (?,?,?,?)
  `
	statement, err := sqlite.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = statement.Exec(book.Name, book.Description, book.Memo, book.IsReading)
	if err != nil {
		return err
	}

	return nil
}

func (sqlite *SqliteDB) CloseDB() error {
	return sqlite.DB.Close()
}

func (sqlite *SqliteDB) SelectAll(isReading bool) ([]Book, error) {
	query := `
  SELECT * FROM books WHERE isReading = ?
  `
	statement, err := sqlite.DB.Prepare(query)
	if err != nil {
		return []Book{}, err
	}

	rows, err := statement.Query(isReading)
	if err != nil {
		return []Book{}, err
	}

	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		book := Book{}
		err = rows.Scan(
			&book.Id,
			&book.Name,
			&book.Description,
			&book.Memo,
			&book.IsReading)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	return books, nil
}

func (sqlite *SqliteDB) Select(id int) (Book, error) {
	query := `
  SELECT * FROM books WHERE id = ?
  `
	statement, err := sqlite.DB.Prepare(query)
	if err != nil {
		fmt.Println("error preparing the query")
		return Book{}, err
	}

	var book Book
	err = statement.QueryRow(id).Scan(
		&book.Id,
		&book.Name,
		&book.Description,
		&book.Memo,
		&book.IsReading)
	if err != nil {
		fmt.Println("error fetching book data")
		return Book{}, err
	}

	return book, nil
}

func (sqlite *SqliteDB) Update(book Book) error {
	query := `
  UPDATE books
  SET name = ?, description = ?, memo = ?, isReading = ?
  WHERE id = ?
  `
	statement, err := sqlite.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		book.Name,
		book.Description,
		book.Memo,
		book.IsReading,
		book.Id)
	if err != nil {
		return err
	}

	return nil
}

func (sqlite *SqliteDB) Delete(id int) error {
	query := `
  DELETE FROM books WHERE id = ?
  `
	statement, err := sqlite.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
