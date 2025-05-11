package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Album struct {
	ID     uint    `json:"id" gorm:"primarykey;autoIncrement"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Year   uint    `json:"year"`
	Price  float64 `json:"price"`
}

func getAlbums(ctx *gin.Context) {
	var albums []Album
	db.Find(&albums)
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(ctx *gin.Context) {
	var newAlbum Album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}

	db.Create(&newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var album Album

	result := db.First(&album, "id = ?", id)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, album)

}

func putAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var modifyAlbum Album //datos ingresados por el usuario
	var album Album       //album de la base de datos

	result := db.First(&album, "id = ?", id)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
		return
	}

	if err := ctx.BindJSON(&modifyAlbum); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "datos incorrectos"})
		return
	}

	album.Title = modifyAlbum.Title
	album.Artist = modifyAlbum.Artist
	album.Year = modifyAlbum.Year
	album.Price = modifyAlbum.Price

	db.Save(&album)
	ctx.IndentedJSON(http.StatusOK, album)
}

func deleteAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result := db.Delete(&Album{}, "id = ?", id)

	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album no encontrado"})
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "album eliminado"})

}

var db *gorm.DB

func initDB() (*gorm.DB, error) {
	//usuario:password@tcp(ruta:puerto)/baseDeDatos
	dsn := "root:4818841ro@tcp(localhost:3307)/db_vinilos"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Conexion a la base de datos exitosa")
	return db, nil
}

func main() {
	var err error
	db, err = initDB()

	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.PUT("/albums/:id", putAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.Run("localhost:8080")
}
