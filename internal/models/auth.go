package models

type AuthRequest struct {
    Login    string `json:"login" example:"user@example.com"`
    Password string `json:"password" example:"qwerty123"`
}

type AuthResponse struct {
    Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type ErrorResponse struct {
    Code    int    `json:"code" example:"400"`
    Message string `json:"message" example:"Invalid data"`
}
