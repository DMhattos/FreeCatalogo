package routes

import (
	"github.com/DMhattos/FreeCatalogo/app/presentation/http/handlers"
	"github.com/gin-gonic/gin"
)

// SetupCategoryRoute configura as rotas relacionadas a categorias no roteador especificado.
func SetupCategoryRoute(categoryHandler *handlers.CategoryHandler, router *gin.RouterGroup) {
	// Rota para criação de Categoria
	router.POST("/categories", categoryHandler.CreateCategoryHandler)

	// Rotas para operações baseadas no ID da categoria.
	router.GET("/categories/:id", categoryHandler.GetCategoryByIDHandler)
	router.PUT("/categories/:id", categoryHandler.UpdateCategoryHandler)
	router.DELETE("/categories/:id", categoryHandler.DeleteCategoryHandler)

	// Rotas para operações baseadas no nome da categoria.
	router.GET("/categories/name/:name", categoryHandler.GetCategoryByNameHandler)
	router.PUT("/categories/name/:name", categoryHandler.UpdateCategoryHandler)
	router.DELETE("/categories/name/:name", categoryHandler.DeleteCategoryHandler)

}
