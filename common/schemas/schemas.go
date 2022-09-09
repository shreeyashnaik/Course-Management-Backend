package schemas

import (
	"github.com/google/uuid"
	"github.com/shreeyashnaik/Course-Management-Backend/pkg/models"
)

// User related Schemas
type SignupPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Course related schemas
type CreateCoursePayload struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	VideoURL    string   `json:"video_url"`
	Topics      []string `json:"topics"`
	Duration    float64  `json:"duration"`
	Category    string   `json:"category"`
	Points      uint     `json:"points"`
}

type ApproveCourseQuery struct {
	CourseID uuid.UUID `json:"course_id"`
}

type DeleteCourseQuery struct {
	CourseID uuid.UUID `json:"course_id"`
}

type ChangeRolePayload struct {
	Role      models.Role `json:"role"`
	ChangeeID uuid.UUID   `json:"changee_id"`
}

type ApprovedCoursesQuery struct {
	Page              int  `json:"page,default=1"`
	PerPage           int  `json:"per_page,default=10"`
	SortByCategoryAsc bool `json:"sort_by_category_asc,default=true"`
}

type ViewCourse struct {
	CourseID          uuid.UUID `json:"course_id"`
	CompletedDuration float64   `json:"completed_duration"`
}
