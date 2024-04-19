package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	DB *sql.DB
}

type Book struct {
	Id            int       `json:"id"`
	Title         string    `json:"title"`
	Subtitle      string    `json:"subtitle"`
	Authors       string    `json:"authors"`
	PublishedDate time.Time `json:"published_date"`
	Thumbnail     string    `json:"thumbnail"`
	InfoLink      string    `json:"info_link"`
	Description   string    `json:"description"`
	Memo          string    `json:"memo"`
	IsReading     bool      `json:"isReading"`
}

func (sqlite *SqliteDB) InitializeDB() {
	database, err := sql.Open("sqlite3", "./books.db")

	if err != nil {
		log.Fatal(err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS books (
    id               INTEGER PRIMARY KEY,
    title            TEXT NOT NULL,
    subtitle         TEXT,
    authors          TEXT,
    publication_date TEXT,
    thumbnail        TEXT,
    info_link        TEXT,
    description      Text,
    memo             TEXT,
    isReading        BOOLEAN);
  `
	_, err = database.Exec(query)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
	}

	sqlite.DB = database
}

func (sqlite *SqliteDB) Insert(book Book) error {
	query := `
  INSERT INTO books (title,
                    subtitle,
                    authors,
                    publication_date,
                    thumbnail,
                    info_link,
                    description,
                    memo,
                    isReading) 
  VALUES (?,?,?,?,?,?,?,?,?)
  `
	statement, err := sqlite.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		book.Title,
		book.Subtitle,
		book.Authors,
		book.PublishedDate,
		book.Thumbnail,
		book.InfoLink,
		book.Description,
		book.Memo,
		book.IsReading)

	if err != nil {
		return err
	}

	return nil
}

func (sqlite *SqliteDB) CloseDB() error {
	return sqlite.DB.Close()
}

func (sqlite *SqliteDB) SelectAll(isReading bool) ([]Book, error) {
	query := "SELECT * FROM books"
	if isReading {
		query = query + " WHERE isReading = true"
	}

	rows, err := sqlite.DB.Query(query)
	if err != nil {
		return []Book{}, err
	}

	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		book := Book{}
		var publishedDate string
		err = rows.Scan(
			&book.Id,
			&book.Title,
			&book.Subtitle,
			&book.Authors,
			&publishedDate,
			&book.Thumbnail,
			&book.InfoLink,
			&book.Description,
			&book.Memo,
			&book.IsReading)
		if err != nil {
			log.Fatal(err)
		}

		parsedDate, err := time.Parse("2006-01-02 15:04:05Z07:00", publishedDate)
		if err != nil {
			return []Book{}, err
		}
		book.PublishedDate = parsedDate
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
	var publishedDate string
	err = statement.QueryRow(id).Scan(
		&book.Id,
		&book.Title,
		&book.Subtitle,
		&book.Authors,
		&publishedDate,
		&book.Thumbnail,
		&book.InfoLink,
		&book.Description,
		&book.Memo,
		&book.IsReading)

	if err != nil {
		fmt.Println("error fetching book data")
		return Book{}, err
	}
	parsedDate, err := time.Parse("2006-01-02 15:04:05Z07:00", publishedDate)
	if err != nil {
		return Book{}, err
	}

	book.PublishedDate = parsedDate
	return book, nil
}

func (sqlite *SqliteDB) Update(book Book) error {
	query := `
  UPDATE books
  SET 
    title = ?,
    subtitle = ?,
    authors = ?,
    publication_date = ?,
    thumbnail = ?,
    info_link = ?,
    description = ?,
    memo = ?,
    isReading = ?
  WHERE id = ?
  `
	statement, err := sqlite.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		book.Title,
		&book.Subtitle,
		&book.Authors,
		&book.PublishedDate,
		&book.Thumbnail,
		&book.InfoLink,
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
