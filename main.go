package main

import (
	"database/sql"
	"ejemplo/internal/api/routes"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectionString := "root:root@tcp(localhost:3306)/api-ejemplo"

	// Abre la conexión a la base de datos
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Comprueba la conexión a la base de datos
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	eng := gin.Default()

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	eng.Run()
}
