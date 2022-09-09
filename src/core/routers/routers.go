package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shreeyashnaik/Course-Management-Backend/src/core/controllers"
	"github.com/shreeyashnaik/Course-Management-Backend/src/core/middlewares"
)

func MountRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Signup
	api.Post("/signup", controllers.Signup)

	// Login
	api.Post("/login", controllers.Login)

	// User
	user := api.Group("/user", middlewares.CheckAuth)
	{
		user.Patch("/change_role", controllers.ChangeRole)

		// Course
		course := user.Group("/course")
		{
			course.Post("/create", controllers.CreateCourse)

			course.Patch("/approve", controllers.ApproveCourse)

			course.Delete("/delete", controllers.DeleteCourse)

			course.Patch("/update", controllers.UpdateCourse)

			course.Get("/all", controllers.FetchApprovedCourses)

			course.Post("/view", controllers.ViewCourse)

			course.Get("/pending", controllers.PendingCourses)

			course.Get("/completed", controllers.CompletedCourses)
		}

	}
}
