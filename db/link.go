package db

import (
	"log"
	"time"

	"github.com/nico-mayer/nym/utils"
)

type Link struct {
	ID         int       `json:"id"`
	Short_code string    `json:"short_code"`
	Long_url   string    `json:"long_url"`
	Created_at time.Time `json:"created_at"`
}

func InsertLink(long_url string) (string, error) {
	short_code := utils.EncryptURL(long_url)
	exists := linkExists(short_code)
	if exists {
		log.Println("Entry for this URL already exists")
		return short_code, nil
	}

	query := "INSERT INTO link (long_url, short_code) VALUES ($1, $2)"
	_, err := DB.Exec(query, long_url, short_code)
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

	err := DB.QueryRow(query, short_code).Scan(&link.ID, &link.Long_url, &link.Short_code, &link.Created_at)

	if err != nil {
		return Link{}, err
	}

	return link, nil
}

func linkExists(short_code string) bool {
	query := "SELECT short_code FROM link WHERE short_code = $1"

	_, err := DB.Exec(query, short_code)

	return err != nil
}
