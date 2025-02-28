package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	router := setupRouter()
	albumsString := AlbumSliceToString(albums)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, albumsString, w.Body.String())
}

func TestPostAlbum(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	testAlbum := Album{
		ID:     "4",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  34.56,
	}

	albumJSON, _ := json.Marshal(testAlbum)
	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(string(albumJSON)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, string(albumJSON), w.Body.String())

	// Remove new element from albums
	RemoveElementFromAlbums("4")
}

func TestGetAlbumByID(t *testing.T) {
	router := setupRouter()
	albumsString := AlbumToString(albums[0])

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, albumsString, w.Body.String())
}

func TestDeleteAlbumByID(t *testing.T) {
	router := setupRouter()
	deletedAlbum := albums[0]

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/albums/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.NotContains(t, albums, deletedAlbum)
}
