package category

import "context"

// CategoryRepository define métodos para acessar e persistir dados de categorias.
type CategoryRepository interface {
	// CreateCategory cria uma nova categoria e retorna a categoria criada ou um erro.
	CreateCategory(ctx context.Context, category *Category) (*Category, error)
	// ListCategories lista todas as categorias disponíveis e retorna uma lista de categorias ou um erro.
	ListCategories(ctx context.Context) ([]*Category, error)
	// GetCategoryByID recupera uma categoria pelo seu ID e retorna a categoria ou um erro se a categoria não for encontrada.
	GetCategoryByID(ctx context.Context, id int) (*Category, error)
	// GetCategoryByName recupera uma categoria pelo seu name e retorna a categoria ou um erro se a categoria não for encontrada.
	GetCategoryByName(ctx context.Context, name string) (*Category, error)
	// UpdateCategory atualiza uma categoria existente e retorna a categoria atualizada ou um erro.
	UpdateCategory(ctx context.Context, category *Category) (*Category, error)
	// DeleteCategory exclui uma categoria pelo seu ID e retorna um erro se a exclusão falhar.
	DeleteCategory(ctx context.Context, id int) error
}
