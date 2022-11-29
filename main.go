package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"
)

type PostError string;

func (e PostError) Error() PostError {
	return e
}

type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Pounce", Artist: "John", Price: 56.99},
	{ID: "2", Title: "Pounce", Artist: "Bob", Price: 26.23},
	{ID: "3", Title: "Pounce", Artist: "Jimmy", Price: 15.45},
}


func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:4000")
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Panicln(err)

			err = err.(PostError)
		}
	}()

	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		panic(PostError("Oh, your json is invalid"))
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, alb := range albums {
		if alb.ID == id {
			c.IndentedJSON(http.StatusOK, alb)

			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not fount"})
}
