package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Controller / handler: Code to map the request path to your logic
func GetAlbums(c *gin.Context) {
	// serialise the josn and add it to response
	c.IndentedJSON(http.StatusOK, Albums)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, album := range Albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
}

func AddAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	Albums = append(Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
