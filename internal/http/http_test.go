package http

import (
	"errors"
	"github.com/Nicholas2012/task_tz/internal/http/mocks"
	"github.com/Nicholas2012/task_tz/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:generate go run github.com/vektra/mockery/v2@latest --all --case=underscore --with-expecter

func TestRandom_OK(t *testing.T) {
	srv, usersMock := setup(t)

	user := storage.User{
		Login: "user-unique-login",
		Data:  []byte(`{"image": "1"}`),
	}

	usersMock.EXPECT().Random(mock.Anything).Return(user, nil).Once()

	resp, err := http.Get(srv.URL + "/random")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var data string
	_, err = resp.Body.Read([]byte(data))
	require.NoError(t, err)
	require.Contains(t, `{"image": "1"}`, data)
	require.Contains(t, "user-unique-login", data)
}

func TestRandom_Error(t *testing.T) {
	srv, usersMock := setup(t)

	usersMock.EXPECT().Random(mock.Anything).Return(storage.User{}, errors.New("bad error")).Once()

	resp, err := http.Get(srv.URL + "/random")
	require.NoError(t, err)

	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestList_OK(t *testing.T) {
	srv, usersMock := setup(t)

	users := []storage.User{
		{
			Login: "user-unique-login",
			Data:  []byte(`{"image": "1"}`),
		},
		{
			Login: "user-unique-login-2",
			Data:  []byte(`{"image": "2"}`),
		},
	}

	usersMock.EXPECT().List(mock.Anything).Return(users, nil).Once()

	resp, err := http.Get(srv.URL + "/list")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var data string
	_, err = resp.Body.Read([]byte(data))
	require.NoError(t, err)

	require.Contains(t, `{"image": "1"}`, data)
	require.Contains(t, "user-unique-login", data)
	require.Contains(t, `{"image": "2"}`, data)
	require.Contains(t, "user-unique-login-2", data)
}

func setup(t *testing.T) (*httptest.Server, *mocks.Users) {
	usersMock := mocks.NewUsers(t)

	router := gin.Default()

	handlers := New(usersMock)
	handlers.Register(router)

	srv := httptest.NewServer(router)
	t.Cleanup(srv.Close)

	return srv, usersMock
}
