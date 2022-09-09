package user

import (
	"github.com/google/uuid"
	"github.com/shreeyashnaik/Course-Management-Backend/pkg/models"
	"gorm.io/gorm"
)

type dbRepo struct {
	DB *gorm.DB
}

func (d *dbRepo) CreateUser(u *models.User) (*models.User, error) {
	if err := d.DB.Create(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (d *dbRepo) GetUserByEmail(email string) (*models.User, error) {
	u := models.User{}
	if err := d.DB.Find(&u, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (d *dbRepo) GetUserByID(id uuid.UUID) (*models.User, error) {
	u := models.User{}
	if err := d.DB.First(&u, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (d *dbRepo) ChangeRole(role models.Role, changeeID uuid.UUID) error {
	if err := d.DB.Model(&models.User{}).Where("id = ?", changeeID).Update("role", role).Error; err != nil {
		return err
	}

	return nil
}

func (d *dbRepo) IncrementRewardPoints(userID uuid.UUID, coursePoints uint) error {
	if err := d.DB.Model(&models.User{}).Where("id = ?", userID).Update("reward_points", gorm.Expr("reward_points + ?", coursePoints)).Error; err != nil {
		return err
	}

	return nil
}

type Repository interface {
	CreateUser(u *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	ChangeRole(role models.Role, changeeID uuid.UUID) error
	IncrementRewardPoints(userID uuid.UUID, coursePoints uint) error
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &dbRepo{
		DB: db,
	}
}
