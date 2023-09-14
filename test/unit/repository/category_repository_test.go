package repository_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	model "github.com/DMhattos/FreeCatalogo/app/domain/category"
	"github.com/DMhattos/FreeCatalogo/app/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestCategoryRepository_InsertCategory(t *testing.T) {
	// Crie um banco de dados simulado e um repositório de categoria.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar o banco de dados simulado: %v", err)
	}
	defer db.Close()
	repo := database.NewCategoryRepository(db)

	// Crie um contexto de teste.
	ctx := context.Background()

	// Crie uma categoria de teste.
	category := &model.Category{
		Name: "Categoria de Teste",
	}

	// Espera uma consulta de inserção com sucesso.
	mock.ExpectExec("INSERT INTO categories").WithArgs(category.Name).WillReturnResult(sqlmock.NewResult(1, 1))

	// Execute a função de inserção.
	_, err = repo.CreateCategory(ctx, category)

	// Verifique se não houve erros.
	assert.NoError(t, err)

	// Verifique se todas as expectativas foram satisfeitas.
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCategoryRepository_GetCategoryByID(t *testing.T) {
	// Crie um banco de dados simulado e um repositório de categoria.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar o banco de dados simulado: %v", err)
	}
	defer db.Close()
	repo := database.NewCategoryRepository(db)

	// Crie um contexto de teste.
	ctx := context.Background()

	// ID de categoria de teste.
	categoryID := 1

	// Crie uma linha de resultado simulada para a consulta SQL.
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(categoryID, "Categoria de Teste")

	// Espera uma consulta de seleção com sucesso.
	mock.ExpectQuery("SELECT id, name FROM categories").WithArgs(categoryID).WillReturnRows(rows)

	// Execute a função de consulta.
	category, err := repo.GetCategoryByID(ctx, categoryID)

	// Verifique se não houve erros.
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, categoryID, category.ID)
	assert.Equal(t, "Categoria de Teste", category.Name)

	// Verifique se todas as expectativas foram satisfeitas.
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCategoryRepository_GetCategoryByName(t *testing.T) {
	// Crie um banco de dados simulado e um repositório de categoria.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar o banco de dados simulado: %v", err)
	}
	defer db.Close()
	repo := database.NewCategoryRepository(db)

	// Crie um contexto de teste.
	ctx := context.Background()

	// Nome da categoria de teste.
	categoryName := "Categoria de Teste"

	// Crie uma linha de resultado simulada para a consulta SQL.
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, categoryName)

	// Espera uma consulta de seleção com sucesso.
	mock.ExpectQuery("SELECT id, name FROM categories").WithArgs(categoryName).WillReturnRows(rows)

	// Execute a função de consulta.
	category, err := repo.GetCategoryByName(ctx, categoryName)

	// Verifique se não houve erros.
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, 1, category.ID)
	assert.Equal(t, categoryName, category.Name)

	// Verifique se todas as expectativas foram satisfeitas.
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCategoryRepository_UpdateCategory(t *testing.T) {
	// Crie um banco de dados simulado e um repositório de categoria.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar o banco de dados simulado: %v", err)
	}
	defer db.Close()
	repo := database.NewCategoryRepository(db)

	// Crie um contexto de teste.
	ctx := context.Background()

	// Crie uma categoria de teste com valores atualizados.
	category := &model.Category{
		ID:   1,
		Name: "Categoria Atualizada",
	}

	// Espera uma consulta de atualização com sucesso.
	mock.ExpectExec("UPDATE categories").WithArgs(category.Name, category.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	// Execute a função de atualização.
	err = repo.UpdateCategory(ctx, category)

	// Verifique se não houve erros.
	assert.NoError(t, err)

	// Verifique se todas as expectativas foram satisfeitas.
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCategoryRepository_DeleteCategory(t *testing.T) {
	// Crie um banco de dados simulado e um repositório de categoria.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar o banco de dados simulado: %v", err)
	}
	defer db.Close()
	repo := database.NewCategoryRepository(db)

	// Crie um contexto de teste.
	ctx := context.Background()

	// ID da categoria que será excluída.
	categoryID := 1

	// Espera uma consulta de exclusão com sucesso.
	mock.ExpectExec("DELETE FROM categories").WithArgs(categoryID).WillReturnResult(sqlmock.NewResult(0, 1))

	// Execute a função de exclusão.
	err = repo.DeleteCategory(ctx, categoryID)

	// Verifique se não houve erros.
	assert.NoError(t, err)

	// Verifique se todas as expectativas foram satisfeitas.
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
