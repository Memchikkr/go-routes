package models

type ErrorResponse struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"Invalid credentials"`
}
