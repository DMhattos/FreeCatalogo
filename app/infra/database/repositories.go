package database

import (
	"database/sql"

	"github.com/DMhattos/FreeCatalogo/app/domain/category"
)

type Repositories struct {
	CategoryRepository category.CategoryRepository
}

func NewRepositories(db *sql.DB) Repositories {
	return Repositories{
		CategoryRepository: NewCategoryRepository(db),
	}
}
