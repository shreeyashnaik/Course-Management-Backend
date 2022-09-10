package course

import (
	"github.com/google/uuid"
	"github.com/shreeyashnaik/Course-Management-Backend/common/schemas"
	"github.com/shreeyashnaik/Course-Management-Backend/pkg/models"
)

type courseSvc struct {
	repository Repository
}

type Service interface {
	CreateCourse(c *schemas.CreateCoursePayload, userID uuid.UUID) (*models.Course, error)
	ApproveCourse(courseID uuid.UUID) error
	DeleteCourse(courseID uuid.UUID, userID uuid.UUID) error
	UpdateCourse(c *models.Course, courseID uuid.UUID, userID uuid.UUID) (*models.Course, error)
	FetchApprovedCourses(p *schemas.ApprovedCoursesQuery) ([]models.Course, error)
	GetCourseByID(id uuid.UUID) (*models.Course, error)
	CreateViewedCourse(course *models.ViewedCourses) (*models.ViewedCourses, error)
	FetchPendingCoursesForUser(userID uuid.UUID) ([]models.Course, error)
	FetchCompletedCoursesForUser(userID uuid.UUID) ([]models.Course, error)
}

func (crs *courseSvc) CreateCourse(c *schemas.CreateCoursePayload, userID uuid.UUID) (*models.Course, error) {
	return crs.repository.CreateCourse(c, userID)
}

func (crs *courseSvc) ApproveCourse(courseID uuid.UUID) error {
	return crs.repository.ApproveCourse(courseID)
}

func (crs *courseSvc) DeleteCourse(courseID uuid.UUID, userID uuid.UUID) error {
	return crs.repository.DeleteCourse(courseID, userID)
}

func (crs *courseSvc) UpdateCourse(c *models.Course, courseID uuid.UUID, userID uuid.UUID) (*models.Course, error) {
	return crs.repository.UpdateCourse(c, courseID, userID)
}

func (crs *courseSvc) FetchApprovedCourses(p *schemas.ApprovedCoursesQuery) ([]models.Course, error) {
	return crs.repository.FetchApprovedCourses(p)
}

func (crs *courseSvc) GetCourseByID(id uuid.UUID) (*models.Course, error) {
	return crs.repository.GetCourseByID(id)
}

func (crs *courseSvc) CreateViewedCourse(course *models.ViewedCourses) (*models.ViewedCourses, error) {
	return crs.repository.CreateViewedCourse(course)
}

func (crs *courseSvc) FetchPendingCoursesForUser(userID uuid.UUID) ([]models.Course, error) {
	return crs.repository.FetchPendingCoursesForUser(userID)
}

func (crs *courseSvc) FetchCompletedCoursesForUser(userID uuid.UUID) ([]models.Course, error) {
	return crs.repository.FetchCompletedCoursesForUser(userID)
}

func NewService(r Repository) Service {
	return &courseSvc{
		repository: r,
	}
}
