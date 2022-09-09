package controllers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shreeyashnaik/Course-Management-Backend/common/db"
	"github.com/shreeyashnaik/Course-Management-Backend/common/schemas"
	"github.com/shreeyashnaik/Course-Management-Backend/common/utils"
	"github.com/shreeyashnaik/Course-Management-Backend/common/views"
	"github.com/shreeyashnaik/Course-Management-Backend/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User
func Signup(ctx *fiber.Ctx) error {
	u := new(schemas.SignupPayload)
	if err := ctx.BodyParser(u); err != nil {
		log.Println(err)
		return views.InvalidParams(ctx)
	}

	// Encrypt Password
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		views.InternalServerError(ctx, err)
	}

	newUser, err := db.UserSvc.CreateUser(&models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: string(encryptedPassword),
	})
	if err != nil {
		views.InternalServerError(ctx, err)
	}

	return views.CreatedView(ctx, newUser)
}

func Login(ctx *fiber.Ctx) error {
	u := new(schemas.LoginPayload)
	if err := ctx.BodyParser(u); err != nil {
		log.Println(err)
		return views.InvalidParams(ctx)
	}

	// Fetch User by Email
	currentUser, err := db.UserSvc.GetUserByEmail(u.Email)
	if err != nil {
		views.InternalServerError(ctx, err)
	}

	// Verify user-entered password
	if err := bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(u.Password)); err != nil {
		return views.UnAuthorisedView(ctx)
	}

	token, err := utils.CreateJWTToken(currentUser.ID.String(), currentUser.Email)
	if err != nil {
		views.InternalServerError(ctx, err)
	}

	return views.OkView(ctx, fiber.Map{
		"user":  currentUser,
		"token": token,
	})
}

func ChangeRole(ctx *fiber.Ctx) error {
	p := new(schemas.ChangeRolePayload)

	userID := ctx.Locals("userId").(uuid.UUID)

	changer, err := db.UserSvc.GetUserByID(userID)
	if err != nil {
		return views.InternalServerError(ctx, err)
	}

	if changer.Role != "superadmin" {
		return views.UnAuthorisedView(ctx)
	}

	if err := db.UserSvc.ChangeRole(p.Role, p.ChangeeID); err != nil {
		return views.InternalServerError(ctx, err)
	}

	return views.OkView(ctx, fiber.Map{
		"changee_id": p.ChangeeID,
		"role":       p.Role,
	})
}

// Course
func CreateCourse(ctx *fiber.Ctx) error {
	p := new(schemas.CreateCoursePayload)
	if err := ctx.BodyParser(p); err != nil {
		return views.InvalidParams(ctx)
	}

	userID := ctx.Locals("userId").(uuid.UUID)

	currentUser, err := db.UserSvc.GetUserByID(userID)
	if err != nil {
		return views.InternalServerError(ctx, err)
	}

	if currentUser.Role != "admin" {
		return views.UnAuthorisedView(ctx)
	}

	course, err := db.CourseSvc.CreateCourse(p, userID)
	if err != nil {
		views.InternalServerError(ctx, err)
	}

	return views.CreatedView(ctx, course)
}

// Super Admin Approval for Course
func ApproveCourse(ctx *fiber.Ctx) error {
	p := new(schemas.ApproveCourseQuery)
	if err := ctx.QueryParser(p); err != nil {
		return views.InvalidParams(ctx)
	}

	userID := ctx.Locals("userId").(uuid.UUID)

	currentUser, err := db.UserSvc.GetUserByID(userID)
	if err != nil {
		return views.InternalServerError(ctx, err)
	}

	if currentUser.Role != "superadmin" {
		return views.UnAuthorisedView(ctx)
	}

	if err := db.CourseSvc.ApproveCourse(p.CourseID); err != nil {
		return views.InternalServerError(ctx, err)
	}

	return views.OkView(ctx, fiber.Map{
		"course_id":   p.CourseID,
		"is_approved": true,
	})
}

func DeleteCourse(ctx *fiber.Ctx) error {
	p := new(schemas.DeleteCourseQuery)
	if err := ctx.QueryParser(p); err != nil {
		return views.InvalidParams(ctx)
	}

	userID := ctx.Locals("userId").(uuid.UUID)

	if err := db.CourseSvc.DeleteCourse(p.CourseID, userID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.UnAuthorisedView(ctx)
		}
		return views.InternalServerError(ctx, err)
	}

	return views.OkView(ctx, fiber.Map{
		"course_id":  p.CourseID,
		"is_deleted": true,
	})
}

func UpdateCourse(ctx *fiber.Ctx) error {
	p := new(models.Course)
	if err := ctx.QueryParser(p); err != nil {
		return views.InvalidParams(ctx)
	}

	course := &models.Course{
		Title:       p.Title,
		Description: p.Description,
		VideoURL:    p.VideoURL,
		Topics:      p.Topics,
		Duration:    p.Duration,
		Category:    p.Category,
		Points:      p.Points,
	}

	userID := ctx.Locals("userId").(uuid.UUID)

	course, err := db.CourseSvc.UpdateCourse(course, p.ID, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.UnAuthorisedView(ctx)
		}
		return views.InternalServerError(ctx, err)
	}

	return views.CreatedView(ctx, course)
}

// Controllers Employee
func FetchApprovedCourses(ctx *fiber.Ctx) error {
	p := new(schemas.ApprovedCoursesQuery)
	if err := ctx.QueryParser(p); err != nil {
		return views.InvalidParams(ctx)
	}

	approvedCourses, err := db.CourseSvc.FetchApprovedCourses(p)
	if err != nil {
		views.InternalServerError(ctx, err)
	}

	return views.OkView(ctx, approvedCourses)
}

func ViewCourse(ctx *fiber.Ctx) error {
	p := new(schemas.ViewCourse)
	if err := ctx.BodyParser(p); err != nil {
		return views.InvalidParams(ctx)
	}

	userID := ctx.Locals("userId").(uuid.UUID)

	course, err := db.CourseSvc.GetCourseByID(p.CourseID)
	if err != nil {
		views.InternalServerError(ctx, err)
	}

	viewedCourse := &models.ViewedCourses{
		UserID:   userID,
		CourseID: p.CourseID,
	}

	if p.CompletedDuration >= course.Duration {
		viewedCourse.CompletedDuration = course.Duration
		viewedCourse.IsCompleted = true

		if err := db.UserSvc.IncrementRewardPoints(userID, course.Points); err != nil {
			return views.InternalServerError(ctx, err)
		}
	}

	viewedCourse.LastViewed = time.Now().Unix()

	_, err = db.CourseSvc.CreateViewedCourse(viewedCourse)
	if err != nil {
		views.InternalServerError(ctx, err)
	}

	return views.OkView(ctx, viewedCourse)
}

func PendingCourses(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uuid.UUID)
	pendingCourses, err := db.CourseSvc.FetchPendingCoursesForUser(userID)
	if err != nil {
		return views.InternalServerError(ctx, err)
	}

	return views.OkView(ctx, pendingCourses)
}

func CompletedCourses(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uuid.UUID)
	completedCourses, err := db.CourseSvc.FetchCompletedCoursesForUser(userID)
	if err != nil {
		return views.InternalServerError(ctx, err)
	}

	return views.OkView(ctx, completedCourses)
}
