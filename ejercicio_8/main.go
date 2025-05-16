package main

import (
	"log"
	"vinyl-api/database"
	handlers "vinyl-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos", err)
	}

	router := gin.Default()
	router.GET("/albums", func(ctx *gin.Context) {
		handlers.GetAlbums(ctx, db)
	})

	router.POST("/albums", func(ctx *gin.Context) {
		handlers.PostAlbums(ctx, db)
	})

	router.GET("/albums/:id", func(ctx *gin.Context) {
		handlers.GetAlbumsByID(ctx, db)
	})
	router.PUT("/albums/:id", func(ctx *gin.Context) {
		handlers.PutAlbumByID(ctx, db)
	})

	router.DELETE("/albums/:id", func(ctx *gin.Context) {
		handlers.DeleteAlbumByID(ctx, db)
	})

	router.Run("localhost:8080")
}
