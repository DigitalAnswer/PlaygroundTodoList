package services

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"

	"github.com/DigitalAnswer/PlaygroundTodoList/models"
)

// UserService struct
type UserService struct {
	db *sql.DB
}

// NewUserService constructor
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// Create user
func (s UserService) Create(user *models.User) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hash)

	// save into DB
	_, err = s.db.Exec("INSERT INTO User(user_name, first_name, last_name, password_hash) VALUES(?, ?, ?, ?)", user.UserName, user.FirstName, user.LastName, user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByID func
func (s UserService) GetByID(id int) (*models.User, error) {
	return nil, nil
}
