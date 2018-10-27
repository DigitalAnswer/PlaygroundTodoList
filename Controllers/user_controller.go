package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"

	"github.com/DigitalAnswer/PlaygroundTodoList/helpers"
	"github.com/DigitalAnswer/PlaygroundTodoList/models"
	"github.com/DigitalAnswer/PlaygroundTodoList/services"
)

// UserController struct
type UserController struct {
	userService *services.UserService
}

// NewUserController constructor
func NewUserController(service *services.UserService) (*UserController, error) {
	return &UserController{
		userService: service,
	}, nil
}

// Mount routes
func (c UserController) Mount(r *Router) {
	r.AddRouteFunc("/signin", c.Signin).Methods(http.MethodGet)
	r.AddRouteFunc("/signin", c.Signin).Methods(http.MethodPost)
	r.AddRouteFunc("/signout", c.Signout).Methods(http.MethodPost)
	r.AddRouteFunc("/signup", c.Signup).Methods(http.MethodPost)
}

// Signin func
func (c *UserController) Signin(w http.ResponseWriter, r *http.Request) {
	req := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Error().Err(err).Msg("Unable to decode")
		helpers.FailureFromError(w, http.StatusInternalServerError, err)
		return
	}

	r.Body.Close()
	user, err := c.userService.Authenticate(req)
	if err != nil {
		log.Error().Err(err).Msg("Cannot be auth")
		helpers.FailureFromError(w, http.StatusForbidden, err)
		return
	}

	jwtToken, _ := c.createJWT(user)
	jsonResponse := map[string]interface{}{
		"user":  user,
		"token": jwtToken,
	}

	JSON(w, http.StatusOK, jsonResponse)
}

// Signout func
func (c *UserController) Signout(w http.ResponseWriter, r *http.Request) {
}

// Signup func
func (c *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	req := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Error().Err(err).Msg("Unable to decode")
		helpers.FailureFromError(w, http.StatusInternalServerError, err)
		return
	}

	defer r.Body.Close()
	user, err := c.userService.Create(req)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create new user")
		helpers.FailureFromError(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusCreated, user)
}

func (c *UserController) createJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	})

	mySigningKey := []byte("toto")
	tokenString, err := token.SignedString(mySigningKey)

	return tokenString, err
}
