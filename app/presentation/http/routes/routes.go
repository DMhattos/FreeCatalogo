package routes

import (
	"github.com/DMhattos/FreeCatalogo/app/presentation/http/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(handlers handlers.Handlers, engine *gin.Engine) {
	api := engine.Group("api/v1")
	{
		SetupCategoryRoute(handlers.CategoryHandler, api)
	}
}
