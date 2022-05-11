package errx

import (
	"github.com/refaldyrk/openapi-pzn/model/domain"
	"github.com/refaldyrk/openapi-pzn/model/web"
)

func ModelInit(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.ID,
		Name: category.Name,
	}
}
