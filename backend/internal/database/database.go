package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	DB *sql.DB
}

type Book struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	Authors       string `json:"authors"`
	PublishedDate string `json:"published_date"`
	Thumbnail     string `json:"thumbnail"`
	InfoLink      string `json:"info_link"`
	Description   string `json:"description"`
	Memo          string `json:"memo"`
	IsReading     bool   `json:"isReading"`
	Read          bool   `json:"read"`
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
    isReading        BOOLEAN NOT NULL,
    read             BOOLEAN NOT NULL);
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
            isReading,
            read)
	VALUES (?,?,?,?,?,?,?,?,?,?)
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
    false,
    false)

	if err != nil {
		return err
	}

	return nil
}

func (sqlite *SqliteDB) CloseDB() error {
	return sqlite.DB.Close()
}

func (sqlite *SqliteDB) SelectAll(isReading bool, read bool) ([]Book, error) {
	query := "SELECT * FROM books"

	params := []string{}
	if isReading {
		q := fmt.Sprintf("isReading = %s", strconv.FormatBool(isReading))
		params = append(params, q)
	}
	if read {
		q := fmt.Sprintf("read = %s", strconv.FormatBool(read))
		params = append(params, q)
	}

	if len(params) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(params, " AND "))
	}

  fmt.Println(query)

	rows, err := sqlite.DB.Query(query)
	if err != nil {
		return []Book{}, err
	}

	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		book := Book{}
		err = rows.Scan(
			&book.Id,
			&book.Title,
			&book.Subtitle,
			&book.Authors,
			&book.PublishedDate,
			&book.Thumbnail,
			&book.InfoLink,
			&book.Description,
			&book.Memo,
			&book.IsReading,
      &book.Read)
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
		&book.Title,
		&book.Subtitle,
		&book.Authors,
		&book.PublishedDate,
		&book.Thumbnail,
		&book.InfoLink,
		&book.Description,
		&book.Memo,
		&book.IsReading,
    &book.Read)

	if err != nil {
		fmt.Println("error fetching book data")
		return Book{}, err
	}
	return book, nil
}

func (sqlite *SqliteDB) Update(id int, book Book) error {
	query := `
	UPDATE books
	SET
		description = ?,
		memo = ?,
		isReading = ?,
    read = ?
	WHERE id = ?
  `
	statement, err := sqlite.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		book.Description,
		book.Memo,
		book.IsReading,
		book.Read,
		id)
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
