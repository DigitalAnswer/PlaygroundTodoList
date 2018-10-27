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
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, err
	}

	user.Password = string(hash)

	// save into DB
	_, err = s.db.Exec("INSERT INTO user(user_name, first_name, last_name, password_hash) VALUES(?, ?, ?, ?)", user.UserName, user.FirstName, user.LastName, user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Authenticate user
func (s UserService) Authenticate(u *models.User) (*models.User, error) {
	user := &models.User{}

	err := s.db.QueryRow("SELECT id, user_name, first_name, last_name, password_hash FROM user WHERE user_name=?", u.UserName).Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByID func
func (s UserService) GetByID(id int) (*models.User, error) {
	return nil, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
