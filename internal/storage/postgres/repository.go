package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Nicholas2012/task_tz/internal/storage"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) List(ctx context.Context) ([]storage.User, error) {
	var users []storage.User
	rows, err := r.db.QueryContext(ctx, "SELECT login, data FROM users")
	if err != nil {
		return nil, fmt.Errorf("query users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user storage.User
		err := rows.Scan(&user.Login, &user.Data)
		if err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *Repository) Get(ctx context.Context, login string) (*storage.User, error) {
	var user storage.User
	err := r.db.QueryRowContext(ctx, "SELECT login, data FROM users WHERE login = $1", login).Scan(&user.Login, &user.Data)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) Save(ctx context.Context, user storage.User) error {
	sql := "INSERT INTO users (login, data) VALUES ($1, $2) ON CONFLICT (login) DO UPDATE SET data = $2"
	_, err := r.db.ExecContext(ctx, sql, user.Login, user.Data)
	return err
}
