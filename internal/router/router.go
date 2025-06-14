package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ToDoInventory/internal/handler"
)

func RegisterRoutes(engine *gin.Engine, db *gorm.DB) {

	handlerDI := handler.NewHandler(db)

	engine.GET("/list", handlerDI.GetToDoList)
	engine.POST("/todo", handlerDI.PostToDo)
	engine.GET("/todo/:id", handlerDI.GetListItemById)
	engine.GET("/todo", handlerDI.GetToDoListByParameters)
	engine.DELETE("/todo/:id", handlerDI.DeleteToDoById)

}
