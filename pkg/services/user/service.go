package user

import (
	"github.com/google/uuid"
	"github.com/shreeyashnaik/Course-Management-Backend/pkg/models"
)

type Service interface {
	CreateUser(u *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	ChangeRole(role models.Role, changeeID uuid.UUID) error
	IncrementRewardPoints(userID uuid.UUID, coursePoints uint) error
}

func (u *userSvc) CreateUser(user *models.User) (*models.User, error) {
	return u.repository.CreateUser(user)
}

func (u *userSvc) GetUserByEmail(email string) (*models.User, error) {
	return u.repository.GetUserByEmail(email)
}

func (u *userSvc) GetUserByID(id uuid.UUID) (*models.User, error) {
	return u.repository.GetUserByID(id)
}

func (u *userSvc) ChangeRole(role models.Role, changeeID uuid.UUID) error {
	return u.repository.ChangeRole(role, changeeID)
}

func (u *userSvc) IncrementRewardPoints(userID uuid.UUID, coursePoints uint) error {
	return u.repository.IncrementRewardPoints(userID, coursePoints)
}

type userSvc struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &userSvc{
		repository: r,
	}
}
