package repository

import (
	"context"
	"database/sql"

	"github.com/Alfeenn/api-go/model"
)

type Repository interface {
	FindAll(ctx context.Context, tx *sql.Tx, limit int, offset int) []model.Category
	Find(ctx context.Context, tx *sql.Tx, categoryId int) (model.Category, error)
	Save(ctx context.Context, tx *sql.Tx, category model.Category) model.Category
	Update(ctx context.Context, tx *sql.Tx, category model.Category) model.Category
	Delete(ctx context.Context, tx *sql.Tx, category model.Category)
}
