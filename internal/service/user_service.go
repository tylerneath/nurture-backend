package service

import (
	"context"
	"errors"

	"github.com/tylerneath/nuture-backend/internal/common"
	models "github.com/tylerneath/nuture-backend/internal/model"
	repo "github.com/tylerneath/nuture-backend/internal/repo"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(req common.CreateUserRequest) error
	GetUserByEmail(email string) (*models.User, error)
	AuthenticateUser(req common.LoginUserRequest) error
	RegisterUser(req common.RegisterUserRequest) error
	CreateMessage(req common.CreateMessageRequest) error
}

type userService struct {
	userRepository    repo.UserRepository
	messageRepository repo.MessageRepository
	log               *zap.Logger
}

func NewUserService(ctx context.Context, userRepository repo.UserRepository, messageRepo repo.MessageRepository, log *zap.Logger) UserService {
	return &userService{
		userRepository: userRepository,
		log:            log,
	}
}

func (s *userService) CreateMessage(req common.CreateMessageRequest) error {
	s.log.Info("creating message", zap.String("message", req.Text))
	return nil
}

func (s *userService) RegisterUser(req common.RegisterUserRequest) error {
	var username *string
	if req.Password == "" || req.Email == "" {
		return ErrInvalidRequest
	}

	if req.Username != "" {
		username = &req.Username
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("error generating new passoword", zap.Error(err))
		return NewInternalServerError(err)
	}

	user := &models.User{
		Email:          req.Email,
		HashedPassword: string(hashedPassword),
		Username:       username,
	}

	if err := s.userRepository.Create(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			s.log.Error("user already exists",
				zap.String("email", req.Email),
				zap.String("userId", user.ID.String()),
			)
			return ErrUserAlreadyExists
		} else {
			s.log.Error("error creating user", zap.String("email", req.Email), zap.Error(err))
			return err
		}
	}
	return nil

}

func (s *userService) AuthenticateUser(req common.LoginUserRequest) error {
	if req.Password == "" || req.Email == "" {
		return ErrInvalidRequest
	}
	user, err := s.userRepository.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.log.Error("user not found", zap.String("email", req.Email))
			return ErrUserNotFound
		} else {
			s.log.Error("error getting user", zap.String("email", req.Email), zap.Error(err))
			return err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password)); err != nil {
		s.log.Error("invalid password", zap.String("email", req.Email))
		return ErrInvalidPassword
	}

	return nil
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.log.Error("user not found", zap.String("email", email))
			return nil, ErrUserNotFound
		} else {
			s.log.Error("error getting user", zap.String("email", email), zap.Error(err))
			return nil, err
		}
	}
	return user, nil
}

func (s *userService) CreateUser(req common.CreateUserRequest) error {

	var username *string
	if req.Username != "" {
		username = &req.Username
	}

	user := &models.User{
		Username: username,
		Email:    req.Email,
	}

	err := s.userRepository.Create(user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			s.log.Error("user already exists",
				zap.String("email", req.Email),
				zap.String("userId", user.ID.String()),
			)
			return ErrUserAlreadyExists
		} else {
			s.log.Error("error creating user", zap.String("email", req.Email), zap.Error(err))
			return err
		}
	}
	return nil
}
