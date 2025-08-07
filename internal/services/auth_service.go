package services

import (
	"errors"
	"fmt"
	"lsp-api/internal/config"
	"lsp-api/internal/models"
	"lsp-api/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

type AuthService interface {
	Register(username, email, password string) error
	AdminLogin(email, password string) (string, error)
	AsesorLogin(email, password string) (string, error)
	AsesiLogin(email, password string) (string, error)
	Login(email, password string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

func (s *authService) AdminLogin(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	if user.Role != "admin" {
		return "", errors.New("unauthorized access")
	}

	if !user.ComparePassword(password) {
		return "", errors.New("invalid email or password")
	}

	return s.generateToken(user)
}

func (s *authService) AsesorLogin(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	if user.Role != "Asesor" {
		return "", errors.New("unauthorized access")
	}

	if !user.ComparePassword(password) {
		return "", errors.New("invalid email or password")
	}

	return s.generateToken(user)
}

func (s *authService) AsesiLogin(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	if user.Role != "Asesi" {
		return "", errors.New("unauthorized access")
	}

	if !user.ComparePassword(password) {
		return "", errors.New("invalid email or password")
	}

	return s.generateToken(user)
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	// Verify password
	if err := s.userRepo.VerifyPassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	return s.generateToken(user)
}

func (s *authService) generateToken(user *models.User) (string, error) {
	expiry, err := time.ParseDuration(s.config.JWTExpiry)
	if err != nil {
		return "", fmt.Errorf("invalid JWT expiry time: %w", err)
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
		"role":     user.Role,
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
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.config.JWTSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return token, nil
}

func (s *authService) Register(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create user with default role "Asesi"
	user := models.User{
		Username:  username,
		Email:     email,
		Password:  string(hashedPassword),
		Role:      "Asesi", // Default role
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Only create user record, don't create asesi record
	return s.userRepo.Create(&user)
}

func (u *models.User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
