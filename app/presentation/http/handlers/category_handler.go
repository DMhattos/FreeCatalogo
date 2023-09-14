package handlers

import (
	"net/http"
	"strconv"

	"github.com/DMhattos/FreeCatalogo/app/application/usecase"
	"github.com/DMhattos/FreeCatalogo/app/domain/category"
	"github.com/DMhattos/FreeCatalogo/app/utils"
	"github.com/gin-gonic/gin"
)

// CategoryHandler é uma estrutura que lida com as solicitações relacionadas a categorias.
type CategoryHandler struct {
	usecase usecase.CategoryUsecase
}

// NewCategoryHandler cria uma nova instância de CategoryHandler.
func NewCategoryHandler(usecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		usecase: usecase,
	}
}

// CreateCategoryHandler lida com solicitações POST para criar uma nova categoria.
func (h *CategoryHandler) CreateCategoryHandler(c *gin.Context) {
	var category category.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar o JSON da categoria"})
		return
	}

	createdCategory, err := h.usecase.CreateCategory(c.Request.Context(), &category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar a categoria"})
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}

// GetCategoryByIDHandler lida com solicitações GET para obter uma categoria por ID.
func (h *CategoryHandler) GetCategoryByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da categoria inválido"})
		return
	}

	foundCategory, err := h.usecase.GetCategoryByID(c.Request.Context(), id)
	if err != nil {
		if err == utils.ErrCategoryNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar a categoria"})
		return
	}

	c.JSON(http.StatusOK, foundCategory)
}

// UpdateCategoryHandler lida com solicitações PUT para atualizar uma categoria.
func (h *CategoryHandler) UpdateCategoryHandler(c *gin.Context) {
	var category category.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar o JSON da categoria"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da categoria inválido"})
		return
	}

	category.ID = id

	updatedCategory, err := h.usecase.UpdateCategory(c.Request.Context(), &category)
	if err != nil {
		if err == utils.ErrCategoryNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar a categoria"})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategoryHandler lida com solicitações DELETE para excluir uma categoria por ID.
func (h *CategoryHandler) DeleteCategoryHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da categoria inválido"})
		return
	}

	err = h.usecase.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		if err == utils.ErrCategoryNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir a categoria"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
