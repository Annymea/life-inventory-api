package router

import (
	"github.com/gin-gonic/gin"

	"ToDoInventory/internal/handler"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.GET("/list", handler.GetToDoList)
	engine.POST("/todo", handler.PostToDo)
	engine.GET("/todo/:id", handler.GetListItemById)

}
