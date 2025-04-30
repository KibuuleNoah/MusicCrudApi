package main

import (
	"MusicCrudApi/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load .env vars
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	// initialize controllers
	sc := controllers.SongsController{}

	// define all song routes
	{
		// group song routes
		songsRouter := router.Group("/songs")
		songsRouter.GET("/", sc.GetAllSongs)
		songsRouter.GET("/:id", sc.GetSong)
		songsRouter.POST("/", sc.CreateSong)
		songsRouter.PUT("/:id", sc.UpdateSong)
		songsRouter.DELETE("/:id", sc.DeleteSong)
	}

	router.Run()
}
