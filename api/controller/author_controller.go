package controller

import (
	"context"
	"strconv"
	"time"

	"github.com/LimJiAn/gin-sqlboiler-example/database"
	"github.com/LimJiAn/gin-sqlboiler-example/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func GetAuthors(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	authors, err := models.Authors().All(ctx, database.DB)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, authors)
}

func GetAuthor(c *gin.Context) {
	id := c.Param("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	author, err := models.FindAuthor(ctx, database.DB, authorId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, author)
}

func NewAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBind(&author); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := author.Insert(ctx, database.DB, boil.Infer()); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, author)
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	author, err := models.FindAuthor(ctx, database.DB, authorId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if _, err := author.Delete(ctx, database.DB); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, author)

}

func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	newAuthor := models.Author{}
	if err := c.ShouldBind(&newAuthor); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	author, err := models.FindAuthor(ctx, database.DB, authorId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	author.Name = newAuthor.Name
	if _, err := author.Update(ctx, database.DB, boil.Infer()); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, author)
}
