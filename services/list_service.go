package services

import (
	"database/sql"

	"github.com/DigitalAnswer/PlaygroundTodoList/models"
)

// ListService struct
type ListService struct {
	db *sql.DB
}

// NewListService constructor
func NewListService(db *sql.DB) *ListService {
	return &ListService{
		db: db,
	}
}

// Create user
func (s ListService) Create(list *models.List) (*models.List, error) {
	// save into DB
	return list, nil

	// res, err := s.db.Exec("INSERT INTO list(name, description) VALUES(?, ?)", list.Name, newNullString(list.Description.String))
	// if err != nil {
	// 	return nil, err
	// } else {
	// 	id, err := res.LastInsertId()
	// 	if err != nil {
	// 		println("Error:", err.Error())
	// 	} else {
	// 		list.ID = id
	// 	}
	// }

	// return list, nil
}
