package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func main() {
	config := mysql.Config{
		User:   "root",
		Passwd: "4031861861",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal("db can not open")
	}

	fmt.Println("db opened")

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("db can not connect")
	}

	fmt.Println("db connected!")

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.Run("127.0.0.1:8080")
}

func getAlbumById(c *gin.Context) {
	var id = c.Param("id")
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	var alb album
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, alb)
}

func getAlbums(c *gin.Context) {
	var albums []album
	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "db error"})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var alb album
		if err = rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			continue
		}
		albums = append(albums, alb)
	}

	c.IndentedJSON(http.StatusOK, albums)
}
