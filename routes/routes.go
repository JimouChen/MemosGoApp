package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"memos/controller"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	memosGroup := r.Group("/memos")
	{
		memosGroup.POST("/add_task", controller.AddTaskController)

	}

	return r
}
