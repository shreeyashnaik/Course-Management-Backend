package schemas

import "github.com/go-playground/validator/v10"

type ErrorView struct {
	FailedField string
	Tag         string
	Value       string
}

func (s *SignupPayload) Validate() []*ErrorView {
	var errors []*ErrorView
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorView{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}

func (s *LoginPayload) Validate() []*ErrorView {
	var errors []*ErrorView
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorView{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}

func (s *CreateCoursePayload) Validate() []*ErrorView {
	var errors []*ErrorView
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorView{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}

func (s *ApproveCourseQuery) Validate() []*ErrorView {
	var errors []*ErrorView
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorView{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}

func (s *DeleteCourseQuery) Validate() []*ErrorView {
	var errors []*ErrorView
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorView{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}

func (s *ChangeRolePayload) Validate() []*ErrorView {
	var errors []*ErrorView
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorView{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}

func (s *ApprovedCoursesQuery) Validate() []*ErrorView {
	var errors []*ErrorView
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorView{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}

func (s *ViewCourse) Validate() []*ErrorView {
	var errors []*ErrorView
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorView{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}
