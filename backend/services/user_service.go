package services

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/seeduler/seeduler/models"
	"github.com/seeduler/seeduler/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
	JWTKey         []byte
}

func NewUserService(userRepository *repositories.UserRepository, jwtKey []byte) *UserService {
	return &UserService{UserRepository: userRepository, JWTKey: jwtKey}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	log.Println("Getting all users (in service)")
	return s.UserRepository.GetUsers()
}

func (s *UserService) SaveUsers(users []models.User) error {
	log.Println("Saving users (in service)")
	return s.UserRepository.SaveUsers(users)
}

func (s *UserService) Authenticate(username, password string) (*models.User, error) {
	users, err := s.UserRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}

	return nil, errors.New("invalid credentials")
}

func (s *UserService) GenerateJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"hall_id":  user.HallID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(s.JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
