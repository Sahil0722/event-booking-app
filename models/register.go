package models

type Register struct {
	ID      int64
	EventID int64 `binding:"required"`
	UserID  int64 `binding:"required"`
}
