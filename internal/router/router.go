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
	engine.POST("/entry", handlerDI.PostEntry)
	engine.GET("/entry/:id", handlerDI.GetListItemById)
	engine.GET("/entry", handlerDI.GetEntryListByParameters)
	engine.DELETE("/entry/:id", handlerDI.DeleteEntryById)
	engine.PUT("/entry", handlerDI.UpdateEntry)

}
