package common

type User struct {
	ID       int    `json:"id" binding:"-" db:"id"`
	Name     string `json:"name" binding:"required,min=3,max=64" db:"name"`
	Username string `json:"username" binding:"required,min=3,max=64" db:"username"`
	Email    string `json:"email" binding:"required,email,max=64" db:"email"`
	Password string `json:"password" binding:"required,min=8,max=64" db:"password"`
}

type UserSession struct {
	ID              int    `json:"id" db:"id"`
	UserID          int    `json:"userId" db:"user_id"`
	UserIP          string `json:"userIp" db:"user_ip"`
	RefreshToken    string `json:"refresh_token" db:"refresh_token"`
	RefreshTokenTTL int64  `json:"refresh_token_ttl" db:"refresh_token_ttl"` // in Unix time format
}
