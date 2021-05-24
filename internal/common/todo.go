package common

type Todo struct {
	ID          int    `json:"id" db:"id"`
	UserID      int    `json:"userId" db:"user_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	NotifyDate  int64  `json:"notifyDate" db:"notify_date"`
	Done        bool   `json:"done" db:"done"`
}
