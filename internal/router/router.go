package router

import (
	"LifeInventoryApi/internal/handler/entryhandler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(engine *gin.Engine, db *gorm.DB) {

	entryHandlerDI := entryhandler.NewEntryHandler(db)

	engine.GET("/docu/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := engine.Group("/api")
	{
		v1 := api.Group("/v1")
		v1.GET("/list", entryHandlerDI.GetEntryList)
		v1.POST("/entry", entryHandlerDI.PostEntry)
		v1.GET("/entry/:id", entryHandlerDI.GetListItemById)
		v1.GET("/entry", entryHandlerDI.GetEntryListByParameters)
		v1.DELETE("/entry/:id", entryHandlerDI.DeleteEntryById)
		v1.PUT("/entry", entryHandlerDI.UpdateEntry)
	}

}
