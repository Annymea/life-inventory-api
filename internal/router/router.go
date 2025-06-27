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

	auth := engine.Group("/auth")
	{
		auth.POST("/auth/signup", authHandlerDI.CreateUser)
		auth.POST("/auth/login", authHandlerDI.Login)
	}

	api := engine.Group("/api")
	{
		//middleware wird nur für die folgenden routen gebraucht. bei den anderern holt man sich ja erst das token
		v1 := api.Group("/v1", middleware.CheckAuth(db))
		{
			v1.GET("/list", entryHandlerDI.GetEntryList)
			v1.POST("/entry", entryHandlerDI.PostEntry)
			v1.GET("/entry/:id", entryHandlerDI.GetListItemById)
			v1.GET("/entry", entryHandlerDI.GetEntryListByParameters)
			v1.DELETE("/entry/:id", entryHandlerDI.DeleteEntryById)
			v1.PUT("/entry", entryHandlerDI.UpdateEntry)
		}
	}

}
