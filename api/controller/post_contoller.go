package controller

import (
	"context"
	"strconv"
	"time"

	"github.com/LimJiAn/gin-sqlboiler-exam/database"
	"github.com/LimJiAn/gin-sqlboiler-exam/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func GetPosts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	posts, err := models.Posts().All(ctx, database.DB)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	post, err := models.FindPost(ctx, database.DB, postId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, post)

}

func NewPost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := post.Insert(ctx, database.DB, boil.Infer()); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, post)

}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	post, err := models.FindPost(ctx, database.DB, postId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if _, err := post.Delete(ctx, database.DB); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, post)

}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	newPost := models.Post{}
	if err := c.ShouldBind(&newPost); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	post, err := models.FindPost(ctx, database.DB, postId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	post.Title = newPost.Title
	if _, err := post.Update(ctx, database.DB, boil.Infer()); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, post)

}
