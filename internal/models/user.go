package models

import (
	"database/sql"
	"time"

	"github.com/Memchikkr/go-routes/internal/utils"
)

type User struct {
	ID           int                     `json:"id" db:"id"`
	CreatedAT    time.Time               `json:"created_at" db:"created_at"`
	PingedAT     time.Time               `json:"pinged_at" db:"pinged_at"`
	TgUsername   string                  `json:"tg_username" db:"tg_username"`
	TgID         int                     `json:"tg_id" db:"tg_id"`
	PhoneNumber  sql.NullString `json:"phone_number" db:"phone_number"`
	Name         sql.NullString `json:"name" db:"name"`
	SurName      sql.NullString `json:"surname" db:"surname"`
	UserName     sql.NullString `json:"username" db:"username"`
	AccessToken  sql.NullString `json:"-" db:"access_token"`
	RefreshToken sql.NullString `json:"-" db:"refresh_token"`
	Description  sql.NullString `json:"description" db:"description"`
}

type UserResponse struct {
	ID          int                     `json:"id" example:"1" db:"id"`
	PhoneNumber utils.NullToEmptyString `json:"phone_number" example:"+89009009090" db:"phone_number"`
	Name        utils.NullToEmptyString `json:"name" example:"name" db:"name"`
	SurName     utils.NullToEmptyString `json:"surname" example:"surname" db:"surname"`
	UserName    utils.NullToEmptyString `json:"username" example:"username" db:"username"`
	Description utils.NullToEmptyString `json:"description" example:"description" db:"description"`
}

type UserPutRequest struct {
	PhoneNumber string `json:"phone_number" example:"+89009009090" db:"phone_number"`
	Name        string `json:"name" example:"name" db:"name"`
	SurName     string `json:"surname" example:"surname" db:"surname"`
	UserName    string `json:"username" example:"username" db:"username"`
	Description string `json:"description" example:"description" db:"description"`
}
