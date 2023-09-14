package category

// CategoryRepository define métodos para acessar e persistir dados de categorias.
type CategoryRepository interface {
	// CreateCategory cria uma nova categoria e retorna a categoria criada ou um erro.
	CreateCategory(category *Category) (*Category, error)
	// ListCategories lista todas as categorias disponíveis e retorna uma lista de categorias ou um erro.
	ListCategories() ([]*Category, error)
	// GetCategoryByID recupera uma categoria pelo seu ID e retorna a categoria ou um erro se a categoria não for encontrada.
	GetCategoryByID(id int) (*Category, error)
	// GetCategoryByName recupera uma categoria pelo seu name e retorna a categoria ou um erro se a categoria não for encontrada.
	GetCategoryByName(name string) (*Category, error)
	// UpdateCategory atualiza uma categoria existente e retorna a categoria atualizada ou um erro.
	UpdateCategory(category *Category) (*Category, error)
	// DeleteCategory exclui uma categoria pelo seu ID e retorna um erro se a exclusão falhar.
	DeleteCategory(id int) error
}
