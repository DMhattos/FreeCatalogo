package handlers

import (
	"github.com/DMhattos/FreeCatalogo/app/application/usecase"
)

type Handlers struct {
	CategoryHandler *CategoryHandler
}

func NewHandlers(usecases usecase.Usecases) Handlers {
	return Handlers{
		CategoryHandler: NewCategoryHandler(usecases.CategoryUsecase),
	}
}
