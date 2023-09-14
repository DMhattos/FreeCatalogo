package usecase_test

import (
	"context"
	"testing"

	"github.com/DMhattos/FreeCatalogo/app/application/usecase"
	model "github.com/DMhattos/FreeCatalogo/app/domain/category"
	"github.com/DMhattos/FreeCatalogo/app/utils"
	"github.com/stretchr/testify/assert"
)

var Categories = []*model.Category{
	{ID: 1, Name: "Livro"},
	{ID: 2, Name: "Filme"},
	{ID: 3, Name: "Mangá"},
	{ID: 4, Name: "Anime"},
	{ID: 5, Name: "Série de TV"},
	{ID: 6, Name: "Quadrinhos"},
	{ID: 7, Name: "Música"},
}

type MockCategoryRepository struct {
	Categories []*model.Category
	LastId     int
}

func NewMockCategoryRepository() *MockCategoryRepository {
	return &MockCategoryRepository{
		Categories: Categories,
		LastId:     len(Categories) + 1,
	}
}

func (r *MockCategoryRepository) CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	// Atribua um ID único à nova categoria.
	category.ID = r.LastId
	r.LastId++

	// Adicione a nova categoria à lista de categorias.
	r.Categories = append(r.Categories, category)

	return category, nil
}

func (r *MockCategoryRepository) ListCategories(ctx context.Context) ([]*model.Category, error) {
	return r.Categories, nil
}

func (r *MockCategoryRepository) GetCategoryByID(ctx context.Context, id int) (*model.Category, error) {
	for _, category := range r.Categories {
		if category.ID == id {
			return category, nil
		}
	}

	return nil, utils.ErrCategoryNotFound
}

func (r *MockCategoryRepository) GetCategoryByName(ctx context.Context, name string) (*model.Category, error) {
	for _, category := range r.Categories {
		if category.Name == name {
			return category, nil
		}
	}

	return nil, utils.ErrCategoryNotFound
}

func (r *MockCategoryRepository) UpdateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	for i, existingCategory := range r.Categories {
		if existingCategory.ID == category.ID {
			// Atualize o nome da categoria.
			r.Categories[i].Name = category.Name
			return category, nil
		}
	}

	return nil, utils.ErrCategoryNotFound
}

func (r *MockCategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	for i, category := range r.Categories {
		if category.ID == id {
			// Remova a categoria da lista.
			r.Categories = append(r.Categories[:i], r.Categories[i+1:]...)
			return nil
		}
	}

	return utils.ErrDeleteFailed
}

func TestCategoryUsecase_CreateCategory(t *testing.T) {
	// Crie um contexto de teste.
	ctx := context.Background()

	// Crie um repositório mock.
	mockRepo := NewMockCategoryRepository()

	// Crie um caso de uso de categoria com o repositório mock.
	categoryUsecase := usecase.NewCategoryUsecase(mockRepo)

	// Caso de teste 1: Categoria com nome válido.
	validCategory := &model.Category{Name: "Nova Categoria"}
	createdCategory, err := categoryUsecase.CreateCategory(ctx, validCategory)
	assert.NoError(t, err)            // Verifique se não há erro.
	assert.NotNil(t, createdCategory) // Verifique se a categoria foi criada.

	// Caso de teste 2: Categoria com nome inválido.
	invalidCategory := &model.Category{Name: ""}
	_, err = categoryUsecase.CreateCategory(ctx, invalidCategory)
	assert.Error(t, err) // Verifique se ocorreu um erro de nome de categoria inválido.
}

func TestCategoryUsecase_GetCategoryByID(t *testing.T) {
	// Crie um contexto de teste.
	ctx := context.Background()

	// Crie um repositório mock.
	mockRepo := NewMockCategoryRepository()

	// Crie um caso de uso de categoria com o repositório mock.
	categoryUsecase := usecase.NewCategoryUsecase(mockRepo)

	// Caso de teste 1: ID de categoria existente.
	existingCategoryID := 1
	foundCategory, err := categoryUsecase.GetCategoryByID(ctx, existingCategoryID)
	assert.NoError(t, err)                                // Verifique se não há erro.
	assert.NotNil(t, foundCategory)                       // Verifique se a categoria foi encontrada.
	assert.Equal(t, existingCategoryID, foundCategory.ID) // Verifique se o ID da categoria está correto.

	// Caso de teste 2: ID de categoria inexistente.
	nonExistentCategoryID := 10
	notFoundCategory, err := categoryUsecase.GetCategoryByID(ctx, nonExistentCategoryID)
	assert.Error(t, err)                            // Verifique se ocorreu um erro.
	assert.Equal(t, utils.ErrCategoryNotFound, err) // Verifique se o erro é de categoria não encontrada.
	assert.Nil(t, notFoundCategory)                 // Verifique se a categoria não foi encontrada (é nula).
}

func TestCategoryUsecase_GetCategoryByName(t *testing.T) {
	// Crie um contexto de teste.
	ctx := context.Background()

	// Crie um repositório mock.
	mockRepo := NewMockCategoryRepository()

	// Crie um caso de uso de categoria com o repositório mock.
	categoryUsecase := usecase.NewCategoryUsecase(mockRepo)

	// Caso de teste 1: Nome de categoria existente.
	existingCategoryName := "Livro"
	foundCategory, err := categoryUsecase.GetCategoryByName(ctx, existingCategoryName)
	assert.NoError(t, err)                                    // Verifique se não há erro.
	assert.NotNil(t, foundCategory)                           // Verifique se a categoria foi encontrada.
	assert.Equal(t, existingCategoryName, foundCategory.Name) // Verifique se o nome da categoria está correto.

	// Caso de teste 2: Nome de categoria inexistente.
	nonExistentCategoryName := "Inexistente"
	notFoundCategory, err := categoryUsecase.GetCategoryByName(ctx, nonExistentCategoryName)
	assert.Error(t, err)                            // Verifique se ocorreu um erro.
	assert.Equal(t, utils.ErrCategoryNotFound, err) // Verifique se o erro é de categoria não encontrada.
	assert.Nil(t, notFoundCategory)                 // Verifique se a categoria não foi encontrada (é nula).

	// Caso de teste 3: Nome de categoria vazio.
	emptyCategoryName := ""
	invalidCategory, err := categoryUsecase.GetCategoryByName(ctx, emptyCategoryName)
	assert.Error(t, err)                               // Verifique se ocorreu um erro.
	assert.Equal(t, utils.ErrInvalidCategoryName, err) // Verifique se o erro é de nome de categoria inválido.
	assert.Nil(t, invalidCategory)                     // Verifique se a categoria não foi encontrada (é nula).
}

func TestCategoryUsecase_UpdateCategory(t *testing.T) {
	// Crie um contexto de teste.
	ctx := context.Background()

	// Crie um repositório mock.
	mockRepo := NewMockCategoryRepository()

	// Crie um caso de uso de categoria com o repositório mock.
	categoryUsecase := usecase.NewCategoryUsecase(mockRepo)

	// Crie uma categoria existente para atualização.
	existingCategory := &model.Category{
		ID:   1,
		Name: "Livro",
	}

	// Caso de teste 1: Atualização bem-sucedida.
	updatedCategoryName := "Novo Livro"
	updatedCategory := &model.Category{
		ID:   existingCategory.ID,
		Name: updatedCategoryName,
	}

	updatedCategory, err := categoryUsecase.UpdateCategory(ctx, updatedCategory)
	assert.NoError(t, err)                                     // Verifique se não há erro.
	assert.NotNil(t, updatedCategory)                          // Verifique se a categoria atualizada não é nula.
	assert.Equal(t, updatedCategoryName, updatedCategory.Name) // Verifique se o nome da categoria foi atualizado corretamente.

	// Caso de teste 2: ID de categoria inexistente.
	nonExistentCategory := &model.Category{
		ID:   10, // ID que não existe no repositório mock.
		Name: "Categoria Inexistente",
	}

	_, err = categoryUsecase.UpdateCategory(ctx, nonExistentCategory)
	assert.Error(t, err)                        // Verifique se ocorreu um erro.
	assert.Equal(t, utils.ErrUpdateFailed, err) // Verifique se o erro é de categoria não encontrada.

	// Caso de teste 3: Nome de categoria vazio.
	emptyCategory := &model.Category{
		ID:   existingCategory.ID,
		Name: "",
	}

	_, err = categoryUsecase.UpdateCategory(ctx, emptyCategory)
	assert.Error(t, err)                               // Verifique se ocorreu um erro.
	assert.Equal(t, utils.ErrInvalidCategoryName, err) // Verifique se o erro é de nome de categoria inválido.
}

func TestCategoryUsecase_DeleteCategory(t *testing.T) {
	// Crie um contexto de teste.
	ctx := context.Background()

	// Crie um repositório mock.
	mockRepo := NewMockCategoryRepository()

	// Crie um caso de uso de categoria com o repositório mock.
	categoryUsecase := usecase.NewCategoryUsecase(mockRepo)

	// Caso de teste 1: Exclusão bem-sucedida.
	existingCategoryID := 1
	err := categoryUsecase.DeleteCategory(ctx, existingCategoryID)
	assert.NoError(t, err) // Verifique se não há erro.

	// Verifique se a categoria foi excluída com sucesso, buscando-a por ID.
	deletedCategory, err := categoryUsecase.GetCategoryByID(ctx, existingCategoryID)
	assert.Equal(t, utils.ErrCategoryNotFound, err) // Verifique se a categoria não foi encontrada após a exclusão.
	assert.Nil(t, deletedCategory)

	// Caso de teste 2: Tentativa de exclusão de uma categoria inexistente.
	nonExistentCategoryID := 10 // ID que não existe no repositório mock.
	err = categoryUsecase.DeleteCategory(ctx, nonExistentCategoryID)
	assert.Equal(t, utils.ErrDeleteFailed, err) // Verifique se ocorreu um erro.

	// Verifique se a categoria não foi excluída, buscando-a por ID.
	_, err = categoryUsecase.GetCategoryByID(ctx, nonExistentCategoryID)
	assert.Equal(t, utils.ErrCategoryNotFound, err) // Verifique se a categoria não foi encontrada.

	// Caso de teste 3: Tentativa de exclusão de uma categoria com ID inválido.
	invalidCategoryID := 0 // ID inválido.
	err = categoryUsecase.DeleteCategory(ctx, invalidCategoryID)
	assert.Equal(t, utils.ErrDeleteFailed, err) // Verifique se ocorreu um erro.

	// Verifique se a categoria não foi excluída, buscando-a por ID.
	_, err = categoryUsecase.GetCategoryByID(ctx, invalidCategoryID)
	assert.Equal(t, utils.ErrCategoryNotFound, err) // Verifique se a categoria não foi encontrada.
}
