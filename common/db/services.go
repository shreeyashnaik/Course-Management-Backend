package db

import (
	"github.com/shreeyashnaik/Course-Management-Backend/pkg/services/course"
	"github.com/shreeyashnaik/Course-Management-Backend/pkg/services/user"
)

var (
	UserSvc   user.Service   = nil
	CourseSvc course.Service = nil
)

func InitServices() {
	db = GetDB()

	userRepo := user.NewPostgresRepo(db)
	UserSvc = user.NewService(userRepo)

	courseRepo := course.NewPostgresRepo(db)
	CourseSvc = course.NewService(courseRepo)
}
