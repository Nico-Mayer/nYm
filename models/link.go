package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nico-mayer/nym/db"
)

type Link struct {
	ID      uuid.UUID `json:"id"`
	Short   string    `json:"short"`
	Long    string    `json:"long"`
	Created time.Time `json:"created"`
}

func InsertLink(long string) error {
	query := "INSERT INTO link (id, short, long, created) VALUES ($1, $2, $3, $4)"
	uuid, _ := uuid.NewRandom()
	short := "http://nym.dev/nym/" + uuid.String()

	_, err := db.DB.Exec(query, uuid, short, long, time.Now())

	if err != nil {
		log.Println("Error inserting new link entry")
		return err
	}
	log.Println("Added a new entry to link table")
	return nil
}
