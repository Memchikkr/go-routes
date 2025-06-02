package models

type AuthRequest struct {
	AuthDate  int    `json:"auth_date" example:"1746370731"`
	FirstName string `json:"first_name" example:"Andrey"`
	Hash      string `json:"hash" example:"90ee68ec25e9b34019e..."`
    Id int `json:"id" example:"123456789"`
    LastName string `json:"last_name" example:"Popov"`
    PhotoUrl string `json:"photo_url" example:"https://t.me/i/userpic/..."`
    UserName string `json:"username" example:"User"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type AuthResponse struct {
    AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
    RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
