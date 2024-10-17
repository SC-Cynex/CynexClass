package models

import "time"

type Role struct {
	RoleID      int       `db:"role_id"`
	RoleName    string    `db:"role_name"`
	AccessLevel int       `db:"access_level"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
