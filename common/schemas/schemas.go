package schemas

import (
	"github.com/google/uuid"
	"github.com/shreeyashnaik/Course-Management-Backend/pkg/models"
)

// User related Schemas
type SignupPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Course related schemas
type CreateCoursePayload struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	VideoURL    string   `json:"video_url" validate:"required"`
	Topics      []string `json:"topics" validate:"required"`
	Duration    float64  `json:"duration" validate:"required"`
	Category    string   `json:"category" validate:"required"`
	Points      uint     `json:"points" validate:"required"`
}

type ApproveCourseQuery struct {
	CourseID uuid.UUID `json:"course_id" validate:"required"`
}

type DeleteCourseQuery struct {
	CourseID uuid.UUID `json:"course_id" validate:"required"`
}

type ChangeRolePayload struct {
	Role      models.Role `json:"role" validate:"required"`
	ChangeeID uuid.UUID   `json:"changee_id" validate:"required"`
}

type ApprovedCoursesQuery struct {
	Page              int  `json:"page,default=1"`
	PerPage           int  `json:"per_page,default=10"`
	SortByCategoryAsc bool `json:"sort_by_category_asc,default=true"`
}

type ViewCourse struct {
	CourseID          uuid.UUID `json:"course_id" validate:"required"`
	CompletedDuration float64   `json:"completed_duration" validate:"required"`
}
