package main

import (
	"database/sql"

	"github.com/firdavsich/surl/random"
)

// get url from postgres
func getUrl(db *sql.DB, hash string) (string, error) {
	var url string
	err := db.QueryRow("SELECT url FROM surl WHERE hash = $1", hash).Scan(&url)
	return url, err
}

// add url to postgres
func addUrl(db *sql.DB, url string) (string, error) {

	hash := random.RandString()
	_, err := db.Exec("INSERT INTO surl (url,hash) VALUES ($1,$2)", url, hash)
	if err != nil {
		err = db.QueryRow("SELECT hash FROM surl WHERE url = $1", url).Scan(&hash)
	}
	return baseURL + hash, err
}
