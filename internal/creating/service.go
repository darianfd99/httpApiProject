package service

import (
	"context"

	mooc "github.com/darianfd99/httpApiProject/internal"
	"github.com/darianfd99/httpApiProject/kit/event"
)

//CourseService is the default CourseService interface
//implementation returned by creating.NewCourseService
type CourseService struct {
	courseRepository mooc.CourseRepository
	eventBus         event.Bus
}

//NewCourseService returns the default Service interface implementation.
func NewCourseService(courseRepository mooc.CourseRepository) CourseService {
	return CourseService{
		courseRepository: courseRepository,
	}
}

//Creating implements the creating.CourseService interface.
func (s CourseService) CreateCourse(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}

	if err := s.courseRepository.Save(ctx, course); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, course.Events())
}
