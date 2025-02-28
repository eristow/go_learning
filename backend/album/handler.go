package album

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	albums, err := DbGetAlbums()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting albums: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting albums"})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := DbGetAlbum(id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting album with ID %s: %v\n", id, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error getting album with ID %s", id)})
		return
	}

	c.JSON(http.StatusOK, album)
}

func PostAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Fprintf(os.Stderr, "Error with album JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error with album JSON"})
		return
	}

	createdAlbum, err := DbCreateAlbum(newAlbum)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error adding album: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error adding album"})
		return
	}

	c.JSON(http.StatusCreated, createdAlbum)
}

func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	deletedAlbum, err := DbDeleteAlbum(id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting album: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting album"})
	}

	c.JSON(http.StatusOK, deletedAlbum)
}
