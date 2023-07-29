package models

type LessonStudent struct {
	LessonId  int
	StudentId int
}

func (LessonStudent) TableName() string {
	return "lesson_student"
}
