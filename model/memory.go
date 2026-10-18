package model

type Memory struct {
	ID          int    `json:"id"`
	GroupID     int    `json:"group_id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
