package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Angelina-Grbchh/lab-user-api/internal/handler"
	"github.com/Angelina-Grbchh/lab-user-api/model"
	"github.com/Angelina-Grbchh/lab-user-api/store"
)

func TestCreateUser(t *testing.T) {
	s := handler.NewServer(store.NewUserStore())
	user := model.User{Name: "Alice", Age: 22, Email: "alice@example.com"}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	w := httptest.NewRecorder()
	s.CreateUser(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
