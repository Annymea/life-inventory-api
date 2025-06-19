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

	engine.GET("/list", handlerDI.GetEntryList)
	engine.POST("/todo", handlerDI.PostEntry)
	engine.GET("/todo/:id", handlerDI.GetListItemById)
	engine.GET("/todo", handlerDI.GetEntryListByParameters)
	engine.DELETE("/todo/:id", handlerDI.DeleteEntryById)
	engine.PUT("/todo", handlerDI.UpdateEntry)

}
