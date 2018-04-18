package models

// User struct
type User struct {
	ID        string `json:"id"`
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"-"`
	Salt      string `json:"-"`
	IsDisable bool   `json:"-"`
}
