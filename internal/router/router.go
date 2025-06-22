package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"LifeInventoryApi/internal/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(engine *gin.Engine, db *gorm.DB) {

	handlerDI := handler.NewHandler(db)

	engine.GET("/docu/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := engine.Group("/api")
	{
		v1 := api.Group("/v1")
		v1.GET("/list", handlerDI.GetEntryList)
		v1.POST("/entry", handlerDI.PostEntry)
		v1.GET("/entry/:id", handlerDI.GetListItemById)
		v1.GET("/entry", handlerDI.GetEntryListByParameters)
		v1.DELETE("/entry/:id", handlerDI.DeleteEntryById)
		v1.PUT("/entry", handlerDI.UpdateEntry)
	}

}
