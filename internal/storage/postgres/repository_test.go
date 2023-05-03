package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/Nicholas2012/task_tz/internal/storage"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/require"
)

func TestRepository(t *testing.T) {
	_, r := setup(t)

	testUser := storage.User{
		Login: "test",
		Data:  []byte(`{"data": "i am a fake user"}`),
	}

	t.Run("Save", func(t *testing.T) {
		err := r.Save(context.Background(), testUser)
		require.NoError(t, err)
	})

	t.Run("Get", func(t *testing.T) {
		user, err := r.Get(context.Background(), testUser.Login)
		require.NoError(t, err)
		require.Equal(t, testUser.Login, user.Login)
		require.Equal(t, testUser.Data, user.Data)
	})

	t.Run("Update", func(t *testing.T) {
		testUser.Data = []byte(`{"data": "become true user"}`)
		err := r.Save(context.Background(), testUser)
		require.NoError(t, err)
	})

	t.Run("GetAfterUpdate", func(t *testing.T) {
		user, err := r.Get(context.Background(), testUser.Login)
		require.NoError(t, err)
		require.Equal(t, testUser.Login, user.Login)
		require.Equal(t, testUser.Data, user.Data)
	})

	t.Run("List", func(t *testing.T) {
		users, err := r.List(context.Background())
		require.NoError(t, err)
		require.Len(t, users, 1)
		require.Equal(t, testUser.Login, users[0].Login)
		require.Equal(t, testUser.Data, users[0].Data)
	})
}

func setup(t *testing.T) (*sql.DB, *Repository) {
	use := os.Getenv("INTEGRATION_TEST")
	if use == "" {
		t.Skipf("INTEGRATION_TEST is not set, skipping integration test")
	}

	pool, err := dockertest.NewPool("")
	require.NoError(t, err)

	require.NoError(t, pool.Client.Ping())

	// Run postgres
	resource, err := pool.Run("postgres", "latest",
		[]string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_DB=users",
		})
	require.NoError(t, err)

	var db *sql.DB

	// Wait for postgres to be ready
	err = pool.Retry(func() error {
		newDB, err := sql.Open(
			"postgres",
			fmt.Sprintf("postgres://postgres:secret@localhost:%s/users?sslmode=disable", resource.GetPort("5432/tcp")),
		)
		if err != nil {
			return err
		}

		if err := newDB.Ping(); err != nil {
			return err
		}

		db = newDB

		return nil
	})
	require.NoError(t, err)

	// Close resource when test is done
	t.Cleanup(func() {
		err := pool.Purge(resource)
		assert.NoError(t, err)
	})

	// Apply migrations
	require.NoError(t, ApplyMigrations(db))

	return db, New(db)
}
