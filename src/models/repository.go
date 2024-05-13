package models

import (
	"database/sql"
	"log"
	"os"
	"techDocumentation/src/constants"

	_ "github.com/go-sql-driver/mysql"
)

type DocumentsRepository interface {
	CreateDocuments(document Document) error
	GetAllDocuments() ([]Document, error)
	DeleteDocuments(document Document) error
}

type DocumentCreator struct {
	db *sql.DB
}

func NewDocumentCreator() (*DocumentCreator, error) {
	db, err := sql.Open("mysql", "root:"+os.Getenv(constants.DB_PASSWORD)+"@tcp(127.0.0.1:3306)/driveDocuments")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error Pinging the DB", err)
		return nil, err
	}

	return &DocumentCreator{
		db: db,
	}, nil
}

func (d *DocumentCreator) CreateDocuments(document Document) error {
	_, err := d.db.Exec(`INSERT INTO documents (id, file_name, extension, owner_email, visibility)
	 VALUES ($1,$2,$3,$4,$5)`,
		document.id, document.fileName,
		document.extension, document.ownerEmail,
		document.visibility)
	if err != nil {
		return err
	}
	return nil
}

func (d *DocumentCreator) GetAllDocuments() ([]Document, error) {
	row, err := d.db.Query("SELECT * FROM documents")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	documents := []Document{}
	for row.Next() {
		document := Document{}
		err = row.Scan(&document.id, &document.fileName, &document.extension, &document.ownerEmail, &document.visibility)
		if err != nil {
			return nil, err
		}

		documents = append(documents, document)
	}
	if row.Err() != nil {
		return nil, err
	}
	return documents, nil
}

func (d *DocumentCreator) DeleteDocuments(document Document) error {
	_, err := d.db.Exec(`DELETE FROM documents WHERE id = $1`, document.id)
	if err != nil {
		return err
	}
	return nil
}
