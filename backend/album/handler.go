package album

import (
	"fmt"
	"go_learning/db"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var dbRepo Repository = NewDatabaseRepo(db.DBConn)

func GetAllAlbums(c *gin.Context) {
	ctx := c.Request.Context()
	albums, err := dbRepo.GetAlbums(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting albums: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting albums"})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	album, err := dbRepo.GetAlbum(ctx, id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting album with ID %s: %v\n", id, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error getting album with ID %s", id)})
		return
	}

	c.JSON(http.StatusOK, album)
}

func PostAlbum(c *gin.Context) {
	ctx := c.Request.Context()
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Fprintf(os.Stderr, "Error with album JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error with album JSON"})
		return
	}

	createdAlbum, err := dbRepo.CreateAlbum(ctx, newAlbum)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error adding album: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error adding album"})
		return
	}

	c.JSON(http.StatusCreated, createdAlbum)
}

func PutAlbumByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	var newAlbum AlbumResponse

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Fprintf(os.Stderr, "Error with album JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error with album JSON"})
		return
	}

	updatedAlbum, err := dbRepo.UpdateAlbum(ctx, id, newAlbum)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating album: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating album"})
		return
	}

	c.JSON(http.StatusOK, updatedAlbum)
}

func DeleteAlbumByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	deletedAlbum, err := dbRepo.DeleteAlbum(ctx, id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting album: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting album"})
		return
	}

	c.JSON(http.StatusOK, deletedAlbum)
}
