package usecase

import (
	"context"

	"github.com/DMhattos/FreeCatalogo/app/domain/category"
	"github.com/DMhattos/FreeCatalogo/app/utils"
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
	// Verifique se o nome da categoria é válido (não vazio).
	if category.Name == "" {
		return nil, utils.ErrInvalidCategoryName
	}

	// Chame a função do repositório para criar a categoria no banco de dados.
	createdCategory, err := uc.repository.CreateCategory(ctx, category)
	if err != nil {
		return nil, utils.ErrInsertFailed
	}

	return createdCategory, nil
}

// GetCategoryByID obtém uma categoria pelo seu ID.
func (uc *CategoryUsecase) GetCategoryByID(ctx context.Context, id int) (*category.Category, error) {
	// Chame a função do repositório para buscar a categoria por ID no banco de dados.
	foundCategory, err := uc.repository.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Verifique se a categoria foi encontrada. Se não, retorne um erro personalizado.
	if foundCategory == nil {
		return nil, utils.ErrCategoryNotFound
	}

	return foundCategory, nil
}

// GetCategoryByName obtém uma categoria pelo seu name.
func (uc *CategoryUsecase) GetCategoryByName(ctx context.Context, name string) (*category.Category, error) {
	// Verifique se o nome da categoria é válido (não vazio).
	if name == "" {
		return nil, utils.ErrInvalidCategoryName
	}

	// Chame a função do repositório para buscar a categoria por name no banco de dados.
	foundCategory, err := uc.repository.GetCategoryByName(ctx, name)
	if err != nil {
		return nil, err
	}

	// Verifique se a categoria foi encontrada. Se não, retorne um erro personalizado.
	if foundCategory == nil {
		return nil, utils.ErrCategoryNotFound
	}

	return foundCategory, nil
}

// AtualizarCategory atualiza uma categoria existente com base nos dados fornecidos.
func (uc *CategoryUsecase) UpdateCategory(ctx context.Context, category *category.Category) (*category.Category, error) {

	if category.Name == "" {
		return nil, utils.ErrInvalidCategoryName
	}

	// Chame a função do repositório para atualizar a categoria no banco de dados.
	err := uc.repository.UpdateCategory(ctx, category)
	if err != nil {
		return nil, utils.ErrUpdateFailed
	}

	return category, nil
}

// DeleteCategory exclui uma categoria pelo seu ID.
func (uc *CategoryUsecase) DeleteCategory(ctx context.Context, id int) error {
	// Chame a função do repositório para excluir a categoria do banco de dados.
	err := uc.repository.DeleteCategory(ctx, id)
	if err != nil {
		return utils.ErrDeleteFailed
	}

	return nil
}
