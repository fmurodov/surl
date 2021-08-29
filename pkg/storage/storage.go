package storage

import (
	"database/sql"
)

// get url from postgres
func Get(db *sql.DB, hash string) (string, error) {
	var url string
	err := db.QueryRow("SELECT url FROM surl WHERE hash = $1", hash).Scan(&url)
	return url, err
}

// add url to postgres
func Add(db *sql.DB, url string) (string, error) {

	hash := RandString()
	_, err := db.Exec("INSERT INTO surl (url,hash) VALUES ($1,$2)", url, hash)
	if err != nil {
		err = db.QueryRow("SELECT hash FROM surl WHERE url = $1", url).Scan(&hash)
	}
	return hash, err
}
