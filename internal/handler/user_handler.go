package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tylerneath/nuture-backend/internal/common"
	"github.com/tylerneath/nuture-backend/internal/service"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserHandler) LoginUser(c *gin.Context) {
	var loginUser common.LoginUserRequest
	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userService.GetUserByEmail(loginUser.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(loginUser.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user logged in successfully"})

}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var newUser common.CreateUserRequest
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newUser.Password == "" || newUser.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating new passoword"})
		return
	}

	newUser.Password = string(hashedPassword)
	if err := u.userService.CreateUser(newUser); err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("user with email %s already exists", newUser.Email)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
