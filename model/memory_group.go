package model

type MemoryGroup struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
