package main

import (
	"database/sql"
	"net/http"

	"github.com/DigitalAnswer/PlaygroundTodoList/controllers"
	"github.com/DigitalAnswer/PlaygroundTodoList/middleware"
	"github.com/DigitalAnswer/PlaygroundTodoList/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", dbDevSettings.dataSource())
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	router := setupRoutes(db)
	http.ListenAndServe(":8080", router)
}

func setupRoutes(db *sql.DB) *controllers.Router {

	router := controllers.NewRouter()

	router.Use(middleware.NewAuthMiddleware("toto"))

	// UserController
	userS := services.NewUserService(db)
	userC, _ := controllers.NewUserController(userS)
	router.AddController(userC)

	// ListController
	listS := services.NewListService(db)
	listC, _ := controllers.NewListController(listS)
	router.AddController(listC)

	return router
}
