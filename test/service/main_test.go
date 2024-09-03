package service

import (
	"os"
	"testing"
	"time"
	"todo-app--go-gin/domain"
	"todo-app--go-gin/service"
)

var todoService service.ITodoService
var userService service.IUserService

func TestMain(m *testing.M) {
	initialTodos := []domain.Todo{
		{
			Id:          1,
			UserId:      1,
			Title:       "Buy groceries",
			Description: "Purchase fruits, vegetables, and bread",
			IsCompleted: false,
			CreatedAt:   mustParseTime("2024-09-01T10:00:00"),
			UpdatedAt:   mustParseTime("2024-09-01T10:00:00"),
		},
		{
			Id:          2,
			UserId:      1,
			Title:       "Complete assignment",
			Description: "Finish the report for the upcoming meeting",
			IsCompleted: true,
			CreatedAt:   mustParseTime("2024-09-02T09:30:00"),
			UpdatedAt:   mustParseTime("2024-09-02T09:30:00"),
		},
		{
			Id:          3,
			UserId:      1,
			Title:       "Workout session",
			Description: "Attend the gym for a cardio session",
			IsCompleted: false,
			CreatedAt:   mustParseTime("2024-09-03T18:00:00"),
			UpdatedAt:   mustParseTime("2024-09-03T18:00:00"),
		},
		{
			Id:          4,
			UserId:      2,
			Title:       "Read a book",
			Description: "Start reading a new novel",
			IsCompleted: true,
			CreatedAt:   mustParseTime("2024-09-04T20:00:00"),
			UpdatedAt:   mustParseTime("2024-09-04T20:00:00"),
		},
	}
	initialUsers := []domain.User{
		{
			Id:       1,
			Username: "user1",
			Email:    "user1@mail.com",
			Password: "$2a$10$xNG1.Ig/Q7kR5l4PDpMHcOw9/xEd.SJWdo2woUBBSM2MIKFoU9eTe",
		},
		{
			Id:       2,
			Username: "user2",
			Email:    "user2@mail.com",
			Password: "$2a$10$xNG1.Ig/Q7kR5l4PDpMHcOw9/xEd.SJWdo2woUBBSM2MIKFoU9eTe",
		},
		{
			Id:       3,
			Username: "user3",
			Email:    "user3@mail.com",
			Password: "$2a$10$xNG1.Ig/Q7kR5l4PDpMHcOw9/xEd.SJWdo2woUBBSM2MIKFoU9eTe",
		},
		{
			Id:       4,
			Username: "user4",
			Email:    "user4@mail.com",
			Password: "$2a$10$xNG1.Ig/Q7kR5l4PDpMHcOw9/xEd.SJWdo2woUBBSM2MIKFoU9eTe",
		},
	}

	fakeTodoRepository := NewFakeTodoRepository(initialTodos)
	fakeUserRepository := NewFakeUserRepository(initialUsers)
	todoService = service.NewTodoService(fakeTodoRepository)
	userService = service.NewUserService(fakeUserRepository)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func parseTime(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05", timeStr)
}

func mustParseTime(timeStr string) time.Time {
	t, err := parseTime(timeStr)
	if err != nil {
		panic(err)
	}
	return t
}
