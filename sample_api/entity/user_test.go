package entity_test

import (
	"sampleApi/entity"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("CreateUser", func(t *testing.T) {
		newUser := entity.User{ID: 1, Username: "test", Password: "test", Email: "test@test.info"}
		if err := newUser.CreateUser(); err != nil {
			t.Fatalf("failed test %#v", err)
		}
	})

	t.Run("CreateUserWithEmptyUsername", func(t *testing.T) {
		newUser := entity.User{ID: 2, Username: "", Password: "test", Email: "test2@test.info"}
		if err := newUser.CreateUser(); err == nil {
			t.Fatalf("expected an error, got nil")
		}
	})

	t.Run("CreateUserWithEmptyPassword", func(t *testing.T) {
		newUser := entity.User{ID: 3, Username: "test3", Password: "", Email: "test3@test.info"}
		if err := newUser.CreateUser(); err == nil {
			t.Fatalf("expected an error, got nil")
		}
	})

	t.Run("CreateUserWithEmptyEmail", func(t *testing.T) {
		newUser := entity.User{ID: 4, Username: "test4", Password: "test4", Email: ""}
		if err := newUser.CreateUser(); err == nil {
			t.Fatalf("expected an error, got nil")
		}
	})
}
