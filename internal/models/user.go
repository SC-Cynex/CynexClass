package models

import "time"

type User struct {
	UserID    int       `db:"user_id"`
	RoleID    int       `db:"role_id"`
	AddressID int       `db:"address_id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Role      int       `json:"role"`
}
