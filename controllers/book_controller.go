package controllers

import (
	"log"
	"net/http"
	"strconv"
	"web-api-go/database"
	"web-api-go/models"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {

		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id has to be integer",
		})
		return
	}
	db := database.GetDB()

	var book models.Book

	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot find id" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"value": "ok",
	})

	c.JSON(http.StatusOK, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDB()
	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot list books" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {

	db := database.GetDB()
	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot create book",
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDB()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot update book",
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {

		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id has to be integer",
		})
		return
	}
	db := database.GetDB()

	err = db.Delete(&models.Book{}, newId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot delete book" + err.Error(),
		})
		return
	}
	c.Status(http.StatusNoContent)
}
