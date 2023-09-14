package usecase

import "github.com/DMhattos/FreeCatalogo/app/infra/database"

type Usecases struct {
	CategoryUsecase CategoryUsecase
}

func NewUsecases(repositories database.Repositories) Usecases {
	return Usecases{
		CategoryUsecase: *NewCategoryUsecase(repositories.CategoryRepository),
	}
}
