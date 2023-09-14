package database

import (
	"context"
	"database/sql"

	model "github.com/DMhattos/FreeCatalogo/app/domain/category"
	"github.com/DMhattos/FreeCatalogo/app/utils"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

// CreateCategory insere uma nova categoria no banco de dados.
func (repo *CategoryRepository) CreateCategory(ctx context.Context, category *model.Category) (
	*model.Category, error) {
	// Verifique se o contexto foi cancelado ou expirou.
	select {
	case <-ctx.Done():
		// O contexto foi cancelado ou expirou.
		return nil, ctx.Err()
	default:
		// O contexto está ativo, continue com a inserção.
	}

	// Crie uma consulta SQL para inserir a categoria no banco de dados.
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id"

	var insertedCategoryID int
	// Execute a consulta usando o banco de dados com o contexto fornecido e obtenha o ID da categoria inserida.
	err := repo.db.QueryRowContext(ctx, query, category.Name).Scan(&insertedCategoryID)

	if err != nil {
		// Trate erros de inserção, se houver algum.
		return nil, err
	}

	// A inserção foi bem-sucedida, então atualize a categoria com o ID gerado.
	category.ID = insertedCategoryID

	// Retorne a categoria recém-inserida.
	return category, nil
}

func (repo *CategoryRepository) ListCategories(ctx context.Context) ([]*model.Category, error) {
	// Verifique se o contexto foi cancelado ou expirou.
	select {
	case <-ctx.Done():
		// O contexto foi cancelado ou expirou.
		return nil, ctx.Err()
	default:
		// O contexto está ativo, continue com a listagem.
	}

	// Crie uma consulta SQL para listar todas as categorias.
	query := "SELECT id, name FROM categories"

	// Execute a consulta usando o banco de dados com o contexto fornecido.
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		// Trate erros de consulta, se houver algum.
		return nil, err
	}
	defer rows.Close()

	// Crie uma lista para armazenar as categorias.
	var categories []*model.Category

	// Percorra as linhas do resultado da consulta.
	for rows.Next() {
		var category model.Category
		// Escaneie os valores da linha para a estrutura Category.
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			// Trate erros de escaneamento.
			return nil, err
		}
		// Adicione a categoria à lista.
		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		// Trate erros de percorrer as linhas.
		return nil, err
	}

	// Retorne a lista de categorias.
	return categories, nil
}

// GetCategoryByID obtém uma categoria pelo seu ID no banco de dados.
func (repo *CategoryRepository) GetCategoryByID(ctx context.Context, id int) (*model.Category, error) {
	// Verifique se o contexto foi cancelado ou expirou.
	select {
	case <-ctx.Done():
		// O contexto foi cancelado ou expirou.
		return nil, ctx.Err()
	default:
		// O contexto está ativo, continue com a consulta.
	}

	// Crie uma consulta SQL para selecionar a categoria com base no ID.
	query := "SELECT id, name FROM categories WHERE id = $1"

	// Execute a consulta usando o banco de dados com o contexto fornecido.
	row := repo.db.QueryRowContext(ctx, query, id)

	var category model.Category
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		// Verifique se ocorreu um erro que não seja um erro de não encontrado.
		if err != sql.ErrNoRows {
			return nil, err
		}
		// Se não foi encontrado nenhum registro, retorne um erro personalizado.
		return nil, utils.ErrCategoryNotFound
	}

	return &category, nil
}

// GetCategoryByName obtém uma categoria pelo seu name no banco de dados.
func (repo *CategoryRepository) GetCategoryByName(ctx context.Context, name string) (*model.Category, error) {
	// Verifique se o contexto foi cancelado ou expirou.
	select {
	case <-ctx.Done():
		// O contexto foi cancelado ou expirou.
		return nil, ctx.Err()
	default:
		// O contexto está ativo, continue com a consulta.
	}

	// Crie uma consulta SQL para selecionar a categoria com base no nome.
	query := "SELECT id, name FROM categories WHERE name = $1"

	// Execute a consulta usando o banco de dados com o contexto fornecido.
	row := repo.db.QueryRowContext(ctx, query, name)

	var category model.Category
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		// Verifique se ocorreu um erro que não seja um erro de não encontrado.
		if err != sql.ErrNoRows {
			return nil, err
		}
		// Se não foi encontrado nenhum registro, retorne um erro personalizado.
		return nil, utils.ErrCategoryNotFound
	}

	return &category, nil
}

// UpdateCategory atualiza uma categoria no banco de dados.
func (repo *CategoryRepository) UpdateCategory(ctx context.Context, category *model.Category) error {
	// Verifique se o contexto foi cancelado ou expirou.
	select {
	case <-ctx.Done():
		// O contexto foi cancelado ou expirou.
		return ctx.Err()
	default:
		// O contexto está ativo, continue com a atualização.
	}

	// Crie uma consulta SQL para atualizar a categoria no banco de dados.
	query := "UPDATE categories SET name = $1 WHERE id = $2"

	// Execute a consulta usando o banco de dados com o contexto fornecido.
	_, err := repo.db.ExecContext(ctx, query, category.Name, category.ID)

	if err != nil {
		// Trate erros de atualização, se houver algum.
		return utils.ErrUpdateFailed
	}

	return nil
}

// DeleteCategory exclui uma categoria do banco de dados.
func (repo *CategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	// Verifique se o contexto foi cancelado ou expirou.
	select {
	case <-ctx.Done():
		// O contexto foi cancelado ou expirou.
		return ctx.Err()
	default:
		// O contexto está ativo, continue com a exclusão.
	}

	// Crie uma consulta SQL para excluir a categoria com base no ID.
	query := "DELETE FROM categories WHERE id = $1"

	// Execute a consulta usando o banco de dados com o contexto fornecido.
	_, err := repo.db.ExecContext(ctx, query, id)

	if err != nil {
		// Trate erros de exclusão, se houver algum.
		return utils.ErrDeleteFailed
	}

	return nil
}
