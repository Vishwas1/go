package main

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	router := gin.Default()
	// TODO: Route: Logic to prepare a response
	router.GET("/", GetAlbums)
	router.GET("/:id", GetAlbumById)
	router.POST("/", AddAlbums)
	return router
}
