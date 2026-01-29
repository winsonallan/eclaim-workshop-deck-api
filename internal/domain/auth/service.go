package auth

import (
	"eclaim-workshop-deck-api/internal/models"
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

// Register - now returns both tokens
func (s *Service) Register(req RegisterRequest) (*models.User, string, string, error) {
	// Check if user exists
	_, err := s.repo.FindByEmail(req.Email)
	if err == nil {
		return nil, "", "", errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", "", err
	}

	// Create user
	user := &models.User{
		RoleNo:    req.RoleNo,
		UserName:  req.Name,
		UserId:    req.UserId,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedBy: &req.CreatedBy,
	}

	if req.UserProfileNo != 0 {
		user.UserProfileNo = &req.UserProfileNo
	}
	if err := s.repo.Create(user); err != nil {
		return nil, "", "", err
	}

	// Generate access token
	accessToken, err := utils.GenerateToken(user.UserNo, s.jwtSecret)
	if err != nil {
		return nil, "", "", err
	}

	// Generate refresh token
	refreshToken, err := utils.GenerateRefreshToken(user.UserNo, s.jwtSecret)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

// Login - now returns both tokens
func (s *Service) Login(req LoginRequest) (*models.User, string, string, error) {
	// Find user
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	// Generate access token
	accessToken, err := utils.GenerateToken(user.UserNo, s.jwtSecret)
	if err != nil {
		return nil, "", "", err
	}

	// Generate refresh token
	refreshToken, err := utils.GenerateRefreshToken(user.UserNo, s.jwtSecret)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

// NEW: RefreshToken service method
func (s *Service) RefreshToken(req RefreshTokenRequest) (string, string, error) {
	// Validate refresh token
	claims, err := utils.ValidateToken(req.RefreshToken, s.jwtSecret)
	if err != nil {
		return "", "", errors.New("invalid or expired refresh token")
	}

	// Verify user still exists
	_, err = s.repo.FindByUserNo(claims.UserNo)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	// Generate new access token
	newAccessToken, err := utils.GenerateToken(claims.UserNo, s.jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Generate new refresh token (refresh token rotation - more secure)
	newRefreshToken, err := utils.GenerateRefreshToken(claims.UserNo, s.jwtSecret)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

// Keep your other methods as they are...
func (s *Service) GetUserByEmail(req FindByEmailRequest) (*models.User, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("User with that email not found!")
	}
	return user, nil
}

func (s *Service) ChangePassword(req ChangePasswordRequest) (*models.User, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("User with that email not found!")
	}

	if user.UserId != req.Username {
		return nil, errors.New("invalid username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid old password")
	}

	if req.NewPassword != req.ConfirmPassword {
		return nil, errors.New("new password and confirmation do not match")
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedNewPassword)
	user.LastModifiedBy = &user.UserNo

	if err := s.repo.ChangePassword(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) UpdateAccount(req UpdateAccountRequest) (*models.User, error) {
	userNo := req.UserNo

	user, err := s.repo.FindByUserNo(userNo)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.UserNo != user.UserNo {
		return nil, errors.New("unauthorized: you can only update your own account")
	}

	if req.Password != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			return nil, errors.New("invalid old password")
		}
	}

	if req.NewPassword != req.ConfirmPassword {
		return nil, errors.New("new password and confirmation do not match")
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedNewPassword)
	user.LastModifiedBy = &user.UserNo

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Username != "" {
		user.UserId = req.Username
	}

	if err := s.repo.UpdateAccount(user); err != nil {
		return nil, err
	}

	return user, nil
}
