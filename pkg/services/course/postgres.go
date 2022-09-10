package course

import (
	"github.com/google/uuid"
	"github.com/shreeyashnaik/Course-Management-Backend/common/schemas"
	"github.com/shreeyashnaik/Course-Management-Backend/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type dbRepo struct {
	DB *gorm.DB
}

type Repository interface {
	CreateCourse(c *schemas.CreateCoursePayload, userID uuid.UUID) (*models.Course, error)
	ApproveCourse(courseID uuid.UUID) error
	UpdateCourse(u *models.Course, courseID, userID uuid.UUID) (*models.Course, error)
	DeleteCourse(courseID uuid.UUID, userID uuid.UUID) error
	FetchApprovedCourses(p *schemas.ApprovedCoursesQuery) ([]models.Course, error)
	GetCourseByID(id uuid.UUID) (*models.Course, error)
	CreateViewedCourse(course *models.ViewedCourses) (*models.ViewedCourses, error)
	FetchPendingCoursesForUser(userID uuid.UUID) ([]models.Course, error)
	FetchCompletedCoursesForUser(userID uuid.UUID) ([]models.Course, error)
}

func (d *dbRepo) CreateCourse(c *schemas.CreateCoursePayload, userID uuid.UUID) (*models.Course, error) {

	newCourse := models.Course{
		Title:       c.Title,
		Description: c.Description,
		VideoURL:    c.VideoURL,
		Topics:      c.Topics,
		Duration:    c.Duration,
		Category:    c.Category,
		Points:      c.Points,
		UserID:      &userID,
	}

	if err := d.DB.Create(&newCourse).Error; err != nil {
		return nil, err
	}

	return &newCourse, nil
}

func (d *dbRepo) ApproveCourse(courseID uuid.UUID) error {
	if err := d.DB.Model(&models.Course{}).Where("id = ?", courseID).Update("is_approved", true).Error; err != nil {
		return err
	}

	return nil
}

func (d *dbRepo) DeleteCourse(courseID uuid.UUID, userID uuid.UUID) error {
	if err := d.DB.First(&models.Course{}, "id = ? AND user_id = ?", courseID, userID).Error; err != nil {
		return err
	}

	if err := d.DB.Delete(&models.Course{}, courseID).Error; err != nil {
		return err
	}

	return nil
}

func (d *dbRepo) UpdateCourse(c *models.Course, courseID, userID uuid.UUID) (*models.Course, error) {
	if err := d.DB.First(&models.Course{}, "id = ? AND user_id = ?", courseID, userID).Error; err != nil {
		return nil, err
	}

	if err := d.DB.Model(&models.Course{ID: courseID}).Updates(c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (d *dbRepo) FetchApprovedCourses(p *schemas.ApprovedCoursesQuery) ([]models.Course, error) {
	approvedCourses := []models.Course{}

	orderByCategory := ""
	if p.SortByCategoryAsc {
		orderByCategory = "category"
	} else {
		orderByCategory = "category DESC"
	}

	if err := d.DB.Offset((p.Page-1)*p.PerPage).Limit(p.PerPage).Order(orderByCategory).Find(&approvedCourses, "is_approved = ?", true).Error; err != nil {
		return nil, err
	}

	return approvedCourses, nil
}

func (d *dbRepo) GetCourseByID(id uuid.UUID) (*models.Course, error) {
	c := models.Course{}
	if err := d.DB.First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func (d *dbRepo) CreateViewedCourse(course *models.ViewedCourses) (*models.ViewedCourses, error) {

	if err := d.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&course).Error; err != nil {
		return nil, err
	}

	return course, nil
}

func (d *dbRepo) FetchPendingCoursesForUser(userID uuid.UUID) ([]models.Course, error) {
	courses := []models.Course{}

	if err := d.DB.Joins("JOIN viewed_courses ON courses.id = viewed_courses.course_id AND viewed_courses.is_completed = false AND viewed_courses.user_id = ?", userID).
		Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

func (d *dbRepo) FetchCompletedCoursesForUser(userID uuid.UUID) ([]models.Course, error) {
	courses := []models.Course{}

	if err := d.DB.Joins("JOIN viewed_courses ON courses.id = viewed_courses.course_id AND viewed_courses.is_completed = true AND viewed_courses.user_id = ?", userID).
		Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &dbRepo{
		DB: db,
	}
}
