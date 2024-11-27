package models

import "time"

type Course struct {
	ID        int       `db:"id_course"`
	Name      string    `db:"name_course"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
