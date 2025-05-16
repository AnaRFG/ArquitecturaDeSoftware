package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	//usuario:password@tcp(ruta:puerto)/baseDeDatos
	dsn := "root:4818841ro@tcp(localhost:3307)/db_vinilos"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Conexion a la base de datos exitosa")
	return db, nil
}
