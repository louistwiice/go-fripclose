package authentication

import (
	"fmt"

	logger "github.com/rs/zerolog/log"

	"github.com/louistwiice/go/fripclose/domain"
	"github.com/louistwiice/go/fripclose/entity"
	"github.com/louistwiice/go/fripclose/utils"
)

type authservice struct {
	repo domain.UserRepository
}

func NewAuthService(r domain.UserRepository) *authservice {
	return &authservice{
		repo: r,
	}
}

func (s *authservice) Create(u *entity.UserCreateUpdate) error {
	hashedPassword, err := utils.HashString(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return s.repo.Create(u)
}

// Service to update authentication date
func (s *authservice) UpdateAuthenticationDate(u *entity.UserDisplay) error {
	return s.repo.UpdateAuthenticationDate(u)
}

// Retrieve a user
func (s *authservice) GetByID(id string) (*entity.UserDisplay, string, error) {
	u, password, err := s.repo.GetByID(id)
	if err != nil {
		return &entity.UserDisplay{}, "", entity.ErrNotFound
	}
	return u, password, nil
}

func (s *authservice) SearchUser(identifier string) (*entity.UserDisplay, string, error) {
	return s.repo.SearchUser(identifier)
}

func (s *authservice) IsAdminOrHasRight(id string) (bool, error) {
	u, _, err := s.repo.GetByID(id)
	if err != nil {
		return false, err
	}

	if u.IsSuperuser || u.IsStaff {
		return true, nil
	}
	return false, nil
}

func (s *authservice) SetActivationCode(u entity.UserDisplay) (string, error) {
	code, err := utils.GenerateOTP(6)
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("%s_activation_code", u.Username)
	response, err := s.repo.SaveTokenInRedis(key, code)

	if err == nil {
		message := utils.ActivationMail(u.FirstName, u.LastName, code)
		utils.SendMail(
			[]string{u.Email},
			"OTP Activation code",
			message,
		)
	}
	logger.Info().Str("to", u.Email).Str("OTP", code).Msg("ACTIVATION_CODE")

	return response, err
}

func (s *authservice) GetActivationCode(token string) (string, error) {
	key := fmt.Sprintf("%s_activation_code", token)
	return s.repo.GetTokenInRedis(key)
}

func (s *authservice) ActivateUser(username string) error {
	return s.repo.ActivateUser(username)
}
