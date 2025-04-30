package controllers

import (
	"MusicCrudApi/database"
	"MusicCrudApi/models"
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

// Master Song Struct Controller
type SongsController struct{}

// reponds with all songs in the DB
func (r SongsController) GetAllSongs(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"songs": database.Songs,
	})
}

// get song by id
func (r SongsController) GetSong(ctx *gin.Context) {

	var song models.Song

	// bind the uri data to song struct
	if err := ctx.ShouldBindUri(&song); err != nil {
		// reponds the err if any
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	// looks for the requested song by ID
	for _, s := range database.Songs {
		if song.ID == s.ID {
			// if found the returns it
			ctx.JSON(http.StatusOK, gin.H{"song": s})
			return
		}
	}

	// reponds that song not found
	ctx.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("Song with ID %d Not Found", song.ID),
	})

}

// create new song
func (r SongsController) CreateSong(ctx *gin.Context) {
	var song models.Song

	if err := ctx.ShouldBind(&song); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// increment DB last idx
	database.DBIdx++
	song.ID = database.DBIdx

	database.Songs = append(database.Songs, song)

	ctx.JSON(http.StatusCreated, gin.H{"song": song})
}

// deletes song
func (r SongsController) DeleteSong(ctx *gin.Context) {

	var song models.Song

	// bind the uri data to song query struct
	if err := ctx.ShouldBindUri(&song); err != nil {
		// repond the err if any
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	// looks for the requested song by ID
	for i, s := range database.Songs {
		if song.ID == s.ID {

			// if found the delete and repond it
			database.Songs = slices.Delete(database.Songs, i, i+1)

			ctx.JSON(http.StatusOK, gin.H{"song": s})
			return
		}
	}
}

// update song
func (r SongsController) UpdateSong(ctx *gin.Context) {

	var song models.Song

	// bind the uri data to song struct
	if err := ctx.ShouldBindUri(&song); err != nil {
		// repond the err if any
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	if err := ctx.ShouldBind(&song); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// looks for the requested song by ID
	for i, s := range database.Songs {
		if song.ID == s.ID {

			if song.Title != "" {
				s.Title = song.Title
			}

			if song.Artist != "" {
				s.Artist = song.Artist
			}

			if song.Genre != "" {
				s.Genre = song.Genre
			}

			if song.Released != 0 {
				s.Released = song.Released
			}

			database.Songs[i] = s

			ctx.JSON(http.StatusOK, gin.H{"song": s})
			return
		}
	}
}
