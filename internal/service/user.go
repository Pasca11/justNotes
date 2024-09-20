package service

import (
	"fmt"
	"github.com/Pasca11/justNotes/internal/repository"
	"github.com/Pasca11/justNotes/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

var secret = os.Getenv("JWT_SECRET")

type UserService interface {
	Register(user *models.User) error
	Login(user *models.User) (*models.LoginResponse, error)
	GetNotes(id int) ([]models.Note, error)
	CreateNote(id int, note *models.Note) error
	DeleteNote(id int) error
}

type UserServiceImpl struct {
	DB repository.UserRepo
}

func NewUserService(db repository.UserRepo) UserService {
	return &UserServiceImpl{
		DB: db,
	}
}

func (s *UserServiceImpl) Register(user *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("userService.Register.Hash: %w", err)
	}
	user.Password = string(hash)
	return s.DB.CreateUser(user)
}

func (s *UserServiceImpl) Login(user *models.User) (*models.LoginResponse, error) {
	saved, err := s.DB.GetUser(user.Username)
	if err != nil {
		return nil, fmt.Errorf("userService.Login: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(saved.Password), []byte(user.Password))
	if err != nil {
		return nil, fmt.Errorf("userService.Login: invalid password")
	}

	token, err := createToken(saved)
	if err != nil {
		return nil, fmt.Errorf("userService.Login: %w", err)
	}

	return &models.LoginResponse{Token: token}, nil
}

func (s *UserServiceImpl) GetNotes(id int) ([]models.Note, error) {
	notes, err := s.DB.GetNotes(id)
	if err != nil {
		return nil, fmt.Errorf("userService.GetNotes: %w", err)
	}
	return notes, nil
}

func (s *UserServiceImpl) CreateNote(id int, note *models.Note) error {
	err := s.DB.CreateNote(id, note)
	if err != nil {
		return fmt.Errorf("userService.CreateNote: %w", err)
	}
	return nil
}

func (s *UserServiceImpl) DeleteNote(id int) error {
	err := s.DB.DeleteNote(id)
	if err != nil {
		return fmt.Errorf("userService.DeleteNote: %w", err)
	}
	return nil
}

func ValidateToken(encodedToken string) error {
	token, err := getTokenFromString(encodedToken)
	if err != nil {
		return fmt.Errorf("userService.ValidateToken: %w", err)
	}
	if !token.Valid {
		return fmt.Errorf("userService.ValidateToken: invalid token")
	}
	return nil
}

func ExtractUserIdFromToken(tokenStr string) (int, error) {
	token, err := getTokenFromString(tokenStr)
	if err != nil {
		return -1, fmt.Errorf("service.extractUserIdFromToken: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return -1, fmt.Errorf("service.extractUserIdFromToken: invalid claims")
	}
	log.Println("USER ID is ", claims["user_id"], int(claims["user_id"].(float64)))
	return int(claims["user_id"].(float64)), nil
}

func ExtractRoleFromToken(tokenStr string) (string, error) {
	token, err := getTokenFromString(tokenStr)
	if err != nil {
		return "", fmt.Errorf("service.extractUserIdFromToken: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("service.extractUserIdFromToken: invalid claims")
	}

	return claims["role"].(string), nil
}

func createToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"role":    user.Role,
			"exp":     time.Now().Add(time.Hour).Unix(),
		})
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func getTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
