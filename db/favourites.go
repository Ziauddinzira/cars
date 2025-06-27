package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB(path string) {
	var err error
	DB, err = sql.Open("sqlite", path)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS favorites (
		id TEXT PRIMARY KEY
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func AddFavorite(carID string) error {
	_, err := DB.Exec("INSERT OR IGNORE INTO favorites (id) VALUES (?)", carID)
	return err
}

func GetFavorites() ([]string, error) {
	rows, err := DB.Query("SELECT id FROM favorites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		rows.Scan(&id)
		ids = append(ids, id)
	}
	return ids, nil
}

func RemoveFavorite(id string) error {
	_, err := DB.Exec("DELETE FROM favorites WHERE id = ?", id)
	return err
}
