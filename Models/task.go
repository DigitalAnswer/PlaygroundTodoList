package models

// Task struct
type Task struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Order  int    `json:"order"`
}
