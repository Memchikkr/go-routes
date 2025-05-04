package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int    `json:"id" db:"id"`
	CreatedAT    time.Time `json:"created_at" db:"created_at"`
	PingedAT     time.Time `json:"pinged_at" db:"pinged_at"`
	TgUsername   string `json:"tg_username" db:"tg_username"`
	TgID         int `json:"tg_id" db:"tg_id"`
	PhoneNumber  sql.NullString `json:"phone_number" db:"phone_number"`
	Name         sql.NullString `json:"name" db:"name"`
	SurName      sql.NullString `json:"surname" db:"surname"`
	UserName     sql.NullString `json:"username" db:"username"`
	AccessToken  sql.NullString `json:"-" db:"access_token"`
	RefreshToken sql.NullString `json:"-" db:"refresh_token"`
	Description  sql.NullString `json:"description" db:"description"`
}
