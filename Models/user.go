package models

import (
	"encoding/json"
)

// User struct
type User struct {
	ID        int64  `json:"id"`
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Salt      string `json:"-"`
	IsDisable bool   `json:"-"`
}

// MarshalJSON for User
func (usr User) MarshalJSON() ([]byte, error) {
	var tmp struct {
		ID        int64  `json:"id"`
		UserName  string `json:"username"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	tmp.ID = usr.ID
	tmp.UserName = usr.UserName
	tmp.FirstName = usr.FirstName
	tmp.LastName = usr.LastName
	return json.Marshal(&tmp)
}
