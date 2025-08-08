package services

import (
	"errors"
	"fmt"
	"lsp-api/internal/config"
	"lsp-api/internal/models"
	"lsp-api/internal/repositories"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type authService struct {
	userRepo  repositories.UserRepository
	asesiRepo repositories.AsesiRepository
	config    *config.Config
}

func NewAuthService(userRepo repositories.UserRepository, asesiRepo repositories.AsesiRepository, config *config.Config) AuthService {
	return &authService{
		userRepo:  userRepo,
		asesiRepo: asesiRepo,
		config:    config,
	}
}

type AuthService interface {
	Register(username, email, password string) error
	AdminLogin(email, password string) (string, error)
	AsesorLogin(email, password string) (string, error)
	AsesiLogin(email, password string) (string, error)
	Login(email, password string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetUserProfile(userID uint) (*models.User, error)
}

func (s *authService) Register(username, email, password string) error {
	// Check if user already exists
	existingUser, err := s.userRepo.FindByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	// Create user with Asesi role
	user := &models.User{
		Username: username,
		Email:    email,
		Role:     "Asesi",
		IsActive: true,
	}

	// Hash password
	if err := user.HashPassword(password); err != nil {
		return err
	}

	// Create user
	if err := s.userRepo.Create(user); err != nil {
		return err
	}

	// Create Asesi record
	asesi := &models.Asesi{
		UserID:      user.ID,
		NamaLengkap: username,
		Email:       email,
	}

	if err := s.asesiRepo.Create(asesi); err != nil {
		return err
	}

	// Update user's id_related
	user.IDRelated = &asesi.ID
	return s.userRepo.Update(user)
}

func (s *authService) AdminLogin(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	if user.Role != "Admin" {
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

	if !user.ComparePassword(password) {
		return "", errors.New("invalid email or password")
	}

	return s.generateToken(user)
}

func (s *authService) generateToken(user *models.User) (string, error) {
	expiry, _ := strconv.Atoi(s.config.JWTExpiry[:len(s.config.JWTExpiry)-1])
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Duration(expiry) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWTSecret))
}

func (s *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.config.JWTSecret), nil
	})
}

func (s *authService) GetUserProfile(userID uint) (*models.User, error) {
	return s.userRepo.FindByID(userID)
}
