package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/yamato0204/sqlboiler-sample/internal/domain/model"
	"github.com/yamato0204/sqlboiler-sample/internal/domain/repo"
	"github.com/yamato0204/sqlboiler-sample/internal/pkg/database"
)

type userRepository struct {
	db *database.DB
}

// NewUserRepository は UserRepository を作成します
func NewUserRepository(db *database.DB) repo.UserRepository {
	return &userRepository{
		db: db,
	}
}

// GetUserByID はIDからユーザーを取得します
func (r *userRepository) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	err := r.db.GetContext(ctx, &user, `
	SELECT 
		id,
		name,
		email,
		created_at,
		updated_at
	FROM 
		users
	WHERE
		id = ?
	`, userID)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no rows in infra")
		return nil, err

	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser はユーザー情報を更新します

// CreateUser はユーザーを作成します
