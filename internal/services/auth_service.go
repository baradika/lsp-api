package services

import (
	"errors"
	"fmt"
	"time"

	"lsp-api/internal/config"
	"lsp-api/internal/models"
	"lsp-api/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(username, fullName, email, password string) error
	Login(email, password string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type authService struct {
	userRepo repositories.UserRepository
	config   *config.Config
}

func NewAuthService(userRepo repositories.UserRepository, config *config.Config) AuthService {
	return &authService{
		userRepo: userRepo,
		config:   config,
	}
}

func (s *authService) Register(username, fullName, email, password string) error {
	// Check if user already exists
	_, err := s.userRepo.FindByEmail(email)
	if err == nil {
		return errors.New("user with this email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Create new user
	user := &models.User{
		Username: username,
		FullName: fullName,
		Email:    email,
		Password: password,
	}

	return s.userRepo.Create(user)
}

func (s *authService) Login(email, password string) (string, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	// Verify password
	if !user.ComparePassword(password) {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	expiry, err := time.ParseDuration(s.config.JWTExpiry)
	if err != nil {
		return "", fmt.Errorf("invalid JWT expiry time: %w", err)
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(expiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func (s *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.config.JWTSecret), nil
	})
}