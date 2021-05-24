package common

type Todo struct {
	ID          int    `json:"id" binding:"required" db:"id"`
	UserID      int    `json:"userId" binding:"-" db:"user_id"`
	Name        string `json:"name" binding:"required" db:"name"`
	Description string `json:"description" binding:"required" db:"description"`
	NotifyDate  int64  `json:"notifyDate" binding:"required" db:"notify_date"` // in Unix time format
	Done        bool   `json:"done" binding:"required" db:"done"`
}
