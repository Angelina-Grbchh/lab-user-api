package test

import (
	"context"
	"testing"

	"github.com/Angelina-Grbchh/lab-user-api/model"
	"github.com/Angelina-Grbchh/lab-user-api/store"
)

func TestAddUser(t *testing.T) {
	s := store.NewUserStore()
	u := model.User{Name: "Test", Age: 25, Email: "test@mail.com"}

	created, _ := s.AddUser(context.Background(), u)
	if created.ID != 1 {
		t.Errorf("Expected ID 1, got %d", created.ID)
	}
}

func TestGetUser(t *testing.T) {
	s := store.NewUserStore()
	u := model.User{Name: "Bob", Age: 30, Email: "bob@example.com"}
	added, _ := s.AddUser(context.Background(), u)

	got, err := s.GetUser(context.Background(), added.ID)
	if err != nil || got.Name != u.Name {
		t.Errorf("Failed to get user")
	}
}
