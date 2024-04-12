package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tylerneath/nuture-backend/internal/common"
	"github.com/tylerneath/nuture-backend/internal/service"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService service.UserService
	log         *zap.Logger
}

func NewUserHandler(userService service.UserService, log *zap.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		log:         log,
	}

}

func (u *UserHandler) GetUserByID(c *gin.Context) {
	errors.New("implement me")
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	errors.New("implement me")
}

func (u *UserHandler) LoginUser(c *gin.Context) {
	var loginUser common.LoginUserRequest
	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := u.userService.AuthenticateUser(loginUser); err != nil {
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		case errors.Is(err, service.ErrInvalidPassword):
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	// need to generate a token
	tokenString, err := service.GenerateToken(loginUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Token": fmt.Sprintf("Bearer %s", tokenString)})
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var newUser common.RegisterUserRequest
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.userService.RegisterUser(newUser); err != nil {
		switch {
		case errors.Is(err, service.ErrUserAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("user with email %s already exists", newUser.Email)})
		case errors.Is(err, service.ErrInvalidRequest):
			c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	u.log.Info("user created", zap.Any("email", newUser.Email))
	c.JSON(http.StatusCreated, newUser)

}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var newUser common.CreateUserRequest
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.userService.CreateUser(newUser); err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("user with email %s already exists", newUser.Email)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.log.Info("user created", zap.Any("user", newUser))
	c.JSON(http.StatusCreated, newUser)

}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	errors.New("implement me")
}
