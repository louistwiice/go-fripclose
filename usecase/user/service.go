package service_user

import (

	"github.com/louistwiice/go/fripclose/domain"
	"github.com/louistwiice/go/fripclose/entity"
	"github.com/louistwiice/go/fripclose/utils"
)

type userservice struct {
	repo domain.UserRepository
}

func NewUserService(r domain.UserRepository) *userservice {
	return &userservice{
		repo: r,
	}
}

func (s *userservice) List() ([]*entity.UserDisplay, error) {
	return s.repo.List()
}

func (s *userservice) Create(u *entity.UserCreateUpdate) error {
	hashedPassword, err := utils.HashString(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return s.repo.Create(u)
}

// Retrieve a user
func (s *userservice) GetByID(id string) (*entity.UserDisplay, string, error) {
	u, password, err := s.repo.GetByID(id)
	if err != nil {
		return &entity.UserDisplay{}, "", entity.ErrNotFound
	}
	return u, password, nil
}

func (s *userservice) UpdateUser(u *entity.UserCreateUpdate) error {
	return s.repo.UpdateInfo(u)
}

func (s *userservice) UpdatePassword(u *entity.UserCreateUpdate) error {
	hashed_password, err := utils.HashString(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed_password

	return s.repo.UpdatePassword(u)
}

func (s *userservice) SearchUser(identifier string) (*entity.UserDisplay, string, error) {
	return s.repo.SearchUser(identifier)
}

func (s *userservice) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *userservice) IsAdminOrHasRight(id string) (bool, error) {
	u, _, err := s.repo.GetByID(id)
	if err != nil {
		return false, err
	}

	if (u.IsSuperuser || u.IsStaff) && u.IsActive {
		return true, nil
	}
	return false, nil
}

func (s *userservice) UploadPicture(u *entity.UserDisplay) error {
	return s.repo.UploadPicture(u)
}