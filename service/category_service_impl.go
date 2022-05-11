package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/refaldyrk/openapi-pzn/errx"
	"github.com/refaldyrk/openapi-pzn/model/domain"
	"github.com/refaldyrk/openapi-pzn/model/web"
	"github.com/refaldyrk/openapi-pzn/repository"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
	DB         *sql.DB
	vald       *validator.Validate
}

func (c *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := c.vald.Struct(request)
	errx.ErrorX(err)

	tx, err := c.DB.Begin()
	errx.ErrorX(err)

	defer errx.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = c.Repository.Save(ctx, tx, category)

	return errx.ModelInit(category)
}

func (c *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := c.vald.Struct(request)
	errx.ErrorX(err)

	tx, err := c.DB.Begin()
	errx.ErrorX(err)

	defer errx.CommitOrRollback(tx)

	category, err := c.Repository.FindById(ctx, tx, request.Id)
	errx.ErrorX(err)

	category.Name = request.Name

	category = c.Repository.Update(ctx, tx, category)

	return errx.ModelInit(category)
}

func (c *CategoryServiceImpl) Delete(ctx context.Context, id int) {

	tx, err := c.DB.Begin()
	errx.ErrorX(err)

	defer errx.CommitOrRollback(tx)

	category, err := c.Repository.FindById(ctx, tx, id)
	errx.ErrorX(err)

	c.Repository.Delete(ctx, tx, category)
}

func (c *CategoryServiceImpl) FindById(ctx context.Context, id int) web.CategoryResponse {
	tx, err := c.DB.Begin()
	errx.ErrorX(err)

	defer errx.CommitOrRollback(tx)

	category, err := c.Repository.FindById(ctx, tx, id)
	errx.ErrorX(err)

	return errx.ModelInit(category)
}

func (c *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {

	tx, err := c.DB.Begin()
	errx.ErrorX(err)

	defer errx.CommitOrRollback(tx)

	categories := c.Repository.FindAll(ctx, tx)

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, errx.ModelInit(category))
	}
	return categoryResponses
}
