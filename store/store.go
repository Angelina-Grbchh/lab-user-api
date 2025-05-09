package store

import (
    "context"
    "errors"
    "lab-user-api/model"
    "sync"
)

type UserStore struct {
    mu     sync.Mutex
    users  map[int]model.User
    nextID int
}

func NewUserStore() *UserStore {
    return &UserStore{
        users:  make(map[int]model.User),
        nextID: 1,
    }
}

func (s *UserStore) AddUser(ctx context.Context, user model.User) (model.User, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    user.ID = s.nextID
    s.users[user.ID] = user
    s.nextID++

    return user, nil
}

func (s *UserStore) GetUser(ctx context.Context, id int) (model.User, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    user, exists := s.users[id]
    if !exists {
        return model.User{}, errors.New("not found")
    }
    return user, nil
}

func (s *UserStore) ListUsers(ctx context.Context) ([]model.User, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    users := make([]model.User, 0, len(s.users))
    for _, u := range s.users {
        users = append(users, u)
    }
    return users, nil
}

func (s *UserStore) UpdateUser(ctx context.Context, id int, user model.User) (model.User, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    _, exists := s.users[id]
    if !exists {
        return model.User{}, errors.New("not found")
    }

    user.ID = id
    s.users[id] = user
    return user, nil
}

func (s *UserStore) DeleteUser(ctx context.Context, id int) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    _, exists := s.users[id]
    if !exists {
        return errors.New("not found")
    }

    delete(s.users, id)
    return nil
}
