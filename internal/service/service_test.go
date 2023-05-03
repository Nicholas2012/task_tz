package service

import (
	"context"
	"encoding/json"
	"github.com/Nicholas2012/task_tz/internal/service/mocks"
	"github.com/Nicholas2012/task_tz/internal/storage"
	"github.com/Nicholas2012/task_tz/pkg/randomuser"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:generate go run github.com/vektra/mockery/v2@latest --all --case=underscore --with-expecter

func TestRandom_OK(t *testing.T) {
	ctx, svc := setup(t)

	response := randomuser.Response{
		Results: []randomuser.Result{
			{
				Login: randomuser.Login{
					Username: "user-unique-login",
				},
				Gender: "male",
			},
		},
	}
	b, err := json.Marshal(response.Results[0])
	require.NoError(t, err)

	ctx.cli.EXPECT().Get().Return(&response, nil).Once()
	ctx.repo.EXPECT().Save(mock.Anything, storage.User{Login: "user-unique-login", Data: b}).Return(nil).Once()

	user, err := svc.Random(context.Background())
	require.NoError(t, err)
	require.Equal(t, "user-unique-login", user.Login)
}

func TestRandom_Empty(t *testing.T) {
	ctx, svc := setup(t)

	response := randomuser.Response{
		Results: []randomuser.Result{},
	}

	ctx.cli.EXPECT().Get().Return(&response, nil).Once()

	_, err := svc.Random(context.Background())
	require.EqualError(t, err, "api returns empty results")
}

type testCtx struct {
	repo *mocks.UsersRepository
	cli  *mocks.Random
}

func setup(t *testing.T) (testCtx, *Service) {
	repo := mocks.NewUsersRepository(t)
	cli := mocks.NewRandom(t)

	return testCtx{
		repo: repo,
		cli:  cli,
	}, New(repo, cli)
}
