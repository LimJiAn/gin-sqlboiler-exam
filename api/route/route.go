package route

import (
	"github.com/LimJiAn/gin-sqlboiler-exam/api/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	// Post
	api.GET("/posts", controller.GetPosts)
	api.GET("/posts/:id", controller.GetPost)
	api.POST("/posts", controller.NewPost)
	api.DELETE("/posts/:id", controller.DeletePost)
	api.PUT("/posts/:id", controller.UpdatePost)

	// Author
	api.GET("/authors", controller.GetAuthors)
	api.GET("/authors/:id", controller.GetAuthor)
	api.POST("/authors", controller.NewAuthor)
	api.DELETE("/authors/:id", controller.DeleteAuthor)
	api.PUT("/authors/:id", controller.UpdateAuthor)
}
