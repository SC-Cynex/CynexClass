package models

import "time"

type Address struct {
	ID           int       `db:"id"`
	Street       string    `db:"street"`
	Number       int       `db:"number"`
	Neighborhood string    `db:"neighborhood"`
	City         string    `db:"city"`
	State        string    `db:"state"`
	ZipCode      string    `db:"zip_code"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
