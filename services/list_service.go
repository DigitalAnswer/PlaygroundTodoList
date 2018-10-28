package services

import (
	"database/sql"
	"errors"
	"fmt"

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

// Create list func
func (s ListService) Create(userID int64, list *models.List) (*models.List, error) {
	// save into DB

	if len(list.Name) == 0 {
		return nil, errors.New("Name should not be empty")
	}

	res, err := s.db.Exec("INSERT INTO list(name, description) VALUES(?, ?)", list.Name, newNullString(list.Description.String))
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		println("Error:", err.Error())
		return nil, err
	}
	list.ID = id

	listIndex := 1
	err = s.db.QueryRow("SELECT COUNT(user_id) FROM user_list WHERE user_id=?", userID).Scan(&listIndex)
	if err != nil {
		return nil, err
	}

	fmt.Println("List index: ", listIndex)
	res, err = s.db.Exec("INSERT INTO user_list(user_id, list_id, list_index) VALUES(?, ?, ?)", userID, list.ID, listIndex)
	if err != nil {
		fmt.Println("Remove list: ", list.ID)
		s.db.Exec("DELETE FROM list WHERE id=?", list.ID)
		return nil, err
	}

	return list, nil
}

// GetAllList func
func (s ListService) GetAllList(userID int64) ([]*models.List, error) {

	rows, err := s.db.Query(`
	SELECT user_list.list_id, list.name, list.description, user_list.list_index 
	FROM user_list 
	INNER JOIN list 
	ON user_list.list_id = list.id AND user_list.user_id = ?;`, userID)
	if err != nil {
		return nil, err
	}

	allList := []*models.List{}
	for rows.Next() {
		list := &models.List{}
		err = rows.Scan(&list.ID, &list.Name, &list.Description, &list.ListIndex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		allList = append(allList, list)
	}

	return allList, nil
}

// Delete func
func (s ListService) Delete(userID, listID int64) error {

	_, err := s.db.Exec("DELETE FROM user_list WHERE user_id=? AND list_id=?", userID, listID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = s.db.Exec("DELETE FROM list WHERE id=?", listID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Get List ID func
func (s ListService) Get(userID, listID int64) (*models.List, error) {

	list := &models.List{}

	err := s.db.QueryRow(`
	SELECT list.id, list.name, list.description, user_list.list_index 
	FROM list 
	INNER JOIN user_list 
	ON list.id = ? AND user_list.user_id = ?;`, listID, userID).Scan(&list.ID, &list.Name, &list.Description, &list.ListIndex)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(`
	SELECT id, name, task_index, task_status_id
	FROM task
	WHERE task.list_id = ?
	`, listID)

	for rows.Next() {
		task := &models.Task{}
		err = rows.Scan(&task.ID, &task.Name, &task.TaskIndex, &task.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		list.Tasks = append(list.Tasks, *task)
	}

	return list, nil
}
