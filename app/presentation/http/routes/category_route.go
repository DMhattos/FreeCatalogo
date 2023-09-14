package routes

import (
	"github.com/DMhattos/FreeCatalogo/app/presentation/http/handlers"
	"github.com/gin-gonic/gin"
)

func SetupCategoryRoute(categoryHandler *handlers.CategoryHandler, router *gin.RouterGroup) {

	router.POST("/categories", categoryHandler.CreateCategoryHandler)
	router.GET("/categories/:id", categoryHandler.GetCategoryByIDHandler)
	router.PUT("/categories/:id", categoryHandler.UpdateCategoryHandler)
	router.DELETE("/categories/:id", categoryHandler.DeleteCategoryHandler)
}
