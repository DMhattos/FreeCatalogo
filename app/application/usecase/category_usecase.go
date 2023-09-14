package usecase

import (
	"context"

	"github.com/DMhattos/FreeCatalogo/app/domain/category"
)

type CategoryUsecase struct {
	repository category.CategoryRepository
}

func NewCategoryUsecase(repository category.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{
		repository: repository,
	}
}

// CreateCategory cria uma nova categoria com base nos dados fornecidos.
func (uc *CategoryUsecase) CreateCategory(ctx context.Context, category *category.Category) (*category.Category, error) {
	// Implemente a lógica de criação da categoria aqui.
	// Você pode chamar métodos do repositório, como uc.repository.CreateCategory, dentro deste método.

	return nil, nil
}

// GetCategoryByID obtém uma categoria pelo seu ID.
func (uc *CategoryUsecase) GetCategoryByID(ctx context.Context, id int) (*category.Category, error) {
	// Implemente a lógica de busca da categoria por ID aqui.
	// Você pode chamar métodos do repositório, como uc.repository.GetCategoryByID, dentro deste método.

	return nil, nil
}

// AtualizarCategory atualiza uma categoria existente com base nos dados fornecidos.
func (uc *CategoryUsecase) UpdateCategory(ctx context.Context, category *category.Category) (*category.Category, error) {
	// Implemente a lógica de atualização da categoria aqui.
	// Você pode chamar métodos do repositório, como uc.repository.UpdateCategory, dentro deste método.

	return nil, nil
}

// DeleteCategory exclui uma categoria pelo seu ID.
func (uc *CategoryUsecase) DeleteCategory(ctx context.Context, id int) error {
	// Implemente a lógica de exclusão da categoria aqui.
	// Você pode chamar métodos do repositório, como uc.repository.DeleteCategory, dentro deste método.

	return nil
}
