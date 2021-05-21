package common

type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password"`
}

type UserSession struct {
	ID              int    `json:"id" db:"id"`
	UserID          int    `json:"userId" db:"user_id"`
	RefreshToken    string `json:"refresh_token" db:"refresh_token"`
	RefreshTokenTTL int64  `json:"refresh_token_ttl" db:"refresh_token_ttl"`
}
