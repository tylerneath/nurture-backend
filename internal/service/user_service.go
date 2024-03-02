package service

import (
	"context"
	"errors"

	"github.com/tylerneath/nuture-backend/internal/common"
	models "github.com/tylerneath/nuture-backend/internal/model"
	repo "github.com/tylerneath/nuture-backend/internal/repo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(req common.CreateUserRequest) error
	GetUserByEmail(email string) (*models.User, error)
}

type userService struct {
	userRepository repo.UserRepository
	log            *zap.Logger
}

func NewUserService(ctx context.Context, userRepository repo.UserRepository, log *zap.Logger) UserService {
	return &userService{
		userRepository: userRepository,
		log:            log,
	}
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
