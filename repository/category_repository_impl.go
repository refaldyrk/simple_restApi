package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/refaldyrk/openapi-pzn/errx"
	"github.com/refaldyrk/openapi-pzn/model/domain"
)

type CategoryRepositoryImpl struct {
}

func (r *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into customer (name) values (?)"

	result, err := tx.ExecContext(ctx, SQL, category.Name)
	errx.ErrorX(err)

	id, err := result.LastInsertId()
	errx.ErrorX(err)

	category.ID = int(id)
	return category
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.ID)
	errx.ErrorX(err)

	return category
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.ID)
	errx.ErrorX(err)
}

func (r *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryID int) (domain.Category, error) {
	SQL := "select id, name from category where id = ?"

	row, err := tx.QueryContext(ctx, SQL, categoryID)
	errx.ErrorX(err)

	category := domain.Category{}
	if row.Next() {
		err := row.Scan(&category.ID, &category.Name)
		errx.ErrorX(err)

		return category, nil
	} else {
		return category, errors.New("not found")
	}

}

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"

	rows, err := tx.QueryContext(ctx, SQL)
	errx.ErrorX(err)

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.ID, &category.Name)
		errx.ErrorX(err)

		categories = append(categories, category)
	}

	return categories
}
