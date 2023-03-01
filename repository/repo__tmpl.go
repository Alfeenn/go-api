package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Alfeenn/api-go/helper"
	"github.com/Alfeenn/api-go/model"
)

type RepoTml struct {
}

func NewCategoryRepository() Repository {
	return &RepoTml{}
}

func (repository *RepoTml) FindAll(ctx context.Context, tx *sql.Tx, limit int, offset int) []model.Category {

	SQL := "SELECT *FROM category LIMIT ? OFFSET ?"
	rows, err := tx.Query(SQL, limit, offset)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var categories []model.Category
	for rows.Next() {

		category := model.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories

}
func (Repository *RepoTml) Find(ctx context.Context, tx *sql.Tx, categoryId int) (model.Category, error) {
	SQL := "SELECT id,name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	category := model.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			panic(err)
		}
		return category, nil

	} else {
		return category, errors.New("category not found")
	}
}
func (Repository *RepoTml) Save(ctx context.Context, tx *sql.Tx, category model.Category) model.Category {
	SQL := "INSERT INTO category(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	category.Id = int(id)
	return category
}
func (Repository *RepoTml) Update(ctx context.Context, tx *sql.Tx, category model.Category) model.Category {
	SQL := "UPDATE category SET name = ? WHERE id = ? "
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	if err != nil {
		panic(err)
	}

	return category

}
func (Repository *RepoTml) Delete(ctx context.Context, tx *sql.Tx, category model.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfErr(err)
}
