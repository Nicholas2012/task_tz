package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Nicholas2012/task_tz/internal/storage"
	"github.com/Nicholas2012/task_tz/pkg/randomuser"
)

type UsersRepository interface {
	Get(ctx context.Context, login string) (*storage.User, error)
	Save(ctx context.Context, user storage.User) error
	List(ctx context.Context) ([]storage.User, error)
}

type Random interface {
	Get() (*randomuser.Response, error)
}

type Service struct {
	repo UsersRepository
	cli  Random
}

func New(repo UsersRepository, cli Random) *Service {
	return &Service{
		repo: repo,
		cli:  cli,
	}
}

func (s *Service) Random(ctx context.Context) (storage.User, error) {
	data, err := s.cli.Get()
	if err != nil {
		return storage.User{}, fmt.Errorf("api returns error: %w", err)
	}

	if len(data.Results) == 0 {
		return storage.User{}, fmt.Errorf("api returns empty results")
	}

	userData, err := json.Marshal(data.Results[0])
	if err != nil {
		return storage.User{}, fmt.Errorf("marshal data: %w", err)
	}

	user := storage.User{
		Login: data.Results[0].Login.Username,
		Data:  userData,
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return storage.User{}, fmt.Errorf("save user: %w", err)
	}

	return user, nil
}

func (s *Service) List(ctx context.Context) ([]storage.User, error) {
	return s.repo.List(ctx)
}
