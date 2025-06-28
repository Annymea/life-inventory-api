package router

import (
	"LifeInventoryApi/internal/handler/authhandler"
	"LifeInventoryApi/internal/handler/entryhandler"
	"LifeInventoryApi/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(engine *gin.Engine, db *gorm.DB) {

	entryHandlerDI := entryhandler.NewEntryHandler(db)
	authHandlerDI := authhandler.NewAuthHandler(db)

	engine.GET("/docu/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := engine.Group("/api")
	{
		v1 := api.Group("/v1")

		// Public Auth-Routen (ohne Auth-Middleware)
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authHandlerDI.CreateUser)
			auth.POST("/login", authHandlerDI.Login)
		}

		// Geschützte Routen
		protected := v1.Group("/", middleware.CheckAuth(db))
		{
			protected.GET("/list", entryHandlerDI.GetEntryList)
			protected.POST("/entry", entryHandlerDI.PostEntry)
			protected.GET("/entry/:id", entryHandlerDI.GetListItemById)
			protected.GET("/entry", entryHandlerDI.GetEntryListByParameters)
			protected.DELETE("/entry/:id", entryHandlerDI.DeleteEntryById)
			protected.PUT("/entry", entryHandlerDI.UpdateEntry)
		}
	}

}
