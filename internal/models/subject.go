package models

import "time"

type Subject struct {
	ID        int       `db:"id_subject"`
	Name      string    `db:"name_subject"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
