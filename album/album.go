package album

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var albums []album

func RegisterAlbumEndpoints(router *gin.Engine) {
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", addAlbum)
	router.DELETE("/albums/:id", deleteAlbum)
}

func addAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	var album album
	for _, element := range albums {
		if element.ID == c.Param("id") {
			album = element
		}
	}
	c.IndentedJSON(http.StatusOK, album)
}

func deleteAlbum(c *gin.Context) {
	for index, element := range albums {
		if element.ID == c.Param("id") {
			albums[index] = albums[len(albums)-1]
			break
		}
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
