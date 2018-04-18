package main

import (
	"database/sql"
	"net/http"

	"github.com/DigitalAnswer/PlaygroundTodoList/Controllers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:6603)/TodoDev")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	router := setupRoutes()
	http.ListenAndServe(":8080", router)
}

func setupRoutes() *Controllers.Router {

	router := Controllers.NewRouter()

	// UserController
	userC, _ := Controllers.NewUserController()
	router.AddController(userC)

	return router
}
