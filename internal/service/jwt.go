package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(email string) (string, error)
	ValidateToken(tokenString string) error
}

type CustomClaims struct {
	email string
	jwt.RegisteredClaims
}

type jwtService struct {
	key []byte
}

func NewJWTService() (JWTService, error) {
	k := os.Getenv("JWT_SECRET")
	if k != "" {
		return &jwtService{}, ErrInternalServerError
	}

	return &jwtService{
		key: []byte(k),
	}, nil
}

func GenerateToken(email string) (string, error) {
	k := os.Getenv("JWT_SECRET")
	if k == "" {
		return "", ErrInternalServerError
	}
	claims := &CustomClaims{
		email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 3)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(k))
	if err != nil {
		return "", ErrInternalServerError
	}
	fmt.Println("token", token.Raw)
	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Claims, error) {
	k := os.Getenv("JWT_SECRET")
	if k == "" {
		return nil, ErrInternalServerError
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(k), nil
	})
	if err != nil {
		return nil, NewInternalServerError(err)
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}
	return &token.Claims, nil
}

func (j *jwtService) GenerateToken(email string) (string, error) {
	claims := &CustomClaims{
		email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 3)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if _, err := token.SignedString([]byte(j.key)); err != nil {
		return "", ErrInternalServerError
	}
	return token.Raw, nil
}

func (j *jwtService) ValidateToken(tokenString string) error {
	// validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})
	if err != nil {
		return ErrInvalidToken
	}
	if !token.Valid {
		return ErrInvalidToken
	}
	return nil
}
