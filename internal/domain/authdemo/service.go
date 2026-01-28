package authdemo

import (
	"eclaim-workshop-deck-api/pkg/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *Repository
	jwtSecret string
}

func NewService(repo *Repository, jwtSecret string) *Service {
	return &Service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *Service) Register(req RegisterRequest) (*UserDemo, string, error) {
	// Check if user exists
	_, err := s.repo.FindByEmail(req.Email)
	if err == nil {
		return nil, "", errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	// Create user
	user := &UserDemo{
		Email:    req.Email,
		Password: string(hashedPassword),
		Name:     req.Name,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, "", err
	}

	// Generate token
	token, err := utils.GenerateToken(user.UserNo, s.jwtSecret)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *Service) Login(req LoginRequest) (*UserDemo, string, error) {
	// Find user
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Generate token
	token, err := utils.GenerateToken(user.UserNo, s.jwtSecret)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *Service) GetUserByID(id uint) (*UserDemo, error) {
	return s.repo.FindByID(id)
}
