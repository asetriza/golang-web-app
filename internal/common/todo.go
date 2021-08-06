package common

type Todo struct {
	ID          int    `json:"id" binding:"-" db:"id"`
	UserID      int    `json:"userId" binding:"-" db:"user_id"`
	Name        string `json:"name" binding:"required,min=1,max=255" db:"name"`
	Description string `json:"description" binding:"required,min=1,max=10000" db:"description"`
	NotifyDate  int64  `json:"notifyDate" binding:"required,min=1" db:"notify_date"` // in Unix time format
	Done        bool   `json:"done" db:"done"`
}

func (t *Todo) IsOwner(userID int) bool {
	return t.UserID == userID
}
