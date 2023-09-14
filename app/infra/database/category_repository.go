package database

import (
	"database/sql"

	model "github.com/DMhattos/FreeCatalogo/app/domain/category"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

// InsertCategory insere uma nova categoria no banco de dados.
func (repo *CategoryRepository) InsertCategory(category *model.Category) error {
	// Implemente a lógica de inserção aqui.
	return nil
}

// GetCategoryByID obtém uma categoria pelo seu ID no banco de dados.
func (repo *CategoryRepository) GetCategoryByID(id int) (*model.Category, error) {
	// Implemente a lógica de consulta aqui.
	return nil, nil
}

// GetCategoryByName obtém uma categoria pelo seu name no banco de dados.
func (repo *CategoryRepository) GetCategoryByName(name string) (*model.Category, error) {
	// Implemente a lógica de consulta aqui.
	return nil, nil
}

// UpdateCategory atualiza uma categoria no banco de dados.
func (repo *CategoryRepository) UpdateCategory(category *model.Category) error {
	// Implemente a lógica de atualização aqui.
	return nil
}

// DeleteCategory exclui uma categoria do banco de dados.
func (repo *CategoryRepository) DeleteCategory(id int) error {
	// Implemente a lógica de exclusão aqui.
	return nil
}
