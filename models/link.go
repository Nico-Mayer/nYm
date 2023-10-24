package models

import (
	"crypto/md5"
	"encoding/base64"
	"log"
	"time"

	"github.com/nico-mayer/nym/db"
)

type Link struct {
	ID         int       `json:"id"`
	Short_code string    `json:"short_code"`
	Long_url   string    `json:"long_url"`
	Created_at time.Time `json:"created_at"`
}

func InsertLink(long_url string) (string, error) {
	short_code := createShortCode(long_url)
	exists := isShortCodeExisting(short_code)
	if exists {
		log.Println("Entry for this URL already exists")
		return short_code, nil
	}

	query := "INSERT INTO link (long_url, short_code) VALUES ($1, $2)"
	_, err := db.DB.Exec(query, long_url, short_code)
	if err != nil {
		log.Println("Error inserting new link entry")
		return "", err
	}

	log.Println("Added a new entry to link table")
	return short_code, nil
}

func GetLink(short_code string) (Link, error) {
	var link Link
	query := "SELECT * FROM link WHERE short_code = $1"

	err := db.DB.QueryRow(query, short_code).Scan(&link.ID, &link.Long_url, &link.Short_code, &link.Created_at)

	if err != nil {
		return Link{}, err
	}

	return link, nil
}

func createShortCode(url string) string {
	hash := md5.Sum([]byte(url))
	return base64.URLEncoding.EncodeToString(hash[:])[:6]
}

func isShortCodeExisting(short_code string) bool {
	query := "SELECT short_code FROM link WHERE short_code = $1"

	_, err := db.DB.Exec(query, short_code)

	return err != nil
}
