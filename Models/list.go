package models

// List struct
type List struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description NullString `json:"description"`
	Tasks       []Task     `json:"tasks"`
	Order       int        `json:"order"`
}
