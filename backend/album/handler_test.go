package album

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func setupTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = httptest.NewRequest("GET", "/", nil)
	return ctx, w
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAlbums(ctx context.Context) ([]*AlbumResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*AlbumResponse), args.Error(1)
}

func (m *MockRepository) GetAlbum(ctx context.Context, id string) (*AlbumResponse, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*AlbumResponse), args.Error(1)
}

func (m *MockRepository) CreateAlbum(ctx context.Context, album Album) (*AlbumResponse, error) {
	args := m.Called(ctx, album)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*AlbumResponse), args.Error(1)
}

func (m *MockRepository) UpdateAlbum(ctx context.Context, id string, album AlbumResponse) (*AlbumResponse, error) {
	args := m.Called(ctx, id, album)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*AlbumResponse), args.Error(1)
}

func (m *MockRepository) DeleteAlbum(ctx context.Context, id string) (*AlbumResponse, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*AlbumResponse), args.Error(1)
}

var originalDbRepo *DatabaseRepo

func setupMockRepo() *MockRepository {
	mockRepo := new(MockRepository)
	originalDbRepo = dbRepo.(*DatabaseRepo)
	dbRepo = mockRepo
	return mockRepo
}

func restoreRepo() {
	dbRepo = originalDbRepo
}

type repoWrapper struct {
	mock *MockRepository
}

func (w *repoWrapper) GetAlbums(ctx context.Context) ([]*AlbumResponse, error) {
	return w.mock.GetAlbums(ctx)
}

func (w *repoWrapper) GetAlbum(ctx context.Context, id string) (*AlbumResponse, error) {
	return w.mock.GetAlbum(ctx, id)
}

func (w *repoWrapper) CreateAlbum(ctx context.Context, album Album) (*AlbumResponse, error) {
	return w.mock.CreateAlbum(ctx, album)
}

func (w *repoWrapper) UpdateAlbum(ctx context.Context, id string, album AlbumResponse) (*AlbumResponse, error) {
	return w.mock.UpdateAlbum(ctx, id, album)
}

func (w *repoWrapper) DeleteAlbum(ctx context.Context, id string) (*AlbumResponse, error) {
	return w.mock.DeleteAlbum(ctx, id)
}

func TestGetAllAlbums(t *testing.T) {
	t.Run("successful retrieval", func(t *testing.T) {

		c, w := setupTestContext()

		mockRepo := setupMockRepo()
		defer restoreRepo()

		expectedAlbums := []*AlbumResponse{
			{ID: "1", Title: "Album One", Artist: "Artist One", Price: 9.99, ImageURL: "url1.jpg"},
			{ID: "2", Title: "Album Two", Artist: "Artist Two", Price: 12.99, ImageURL: "url2.jpg"},
		}

		mockRepo.On("GetAlbums", c.Request.Context()).Return(expectedAlbums, nil)

		GetAllAlbums(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseAlbums []*AlbumResponse
		err := json.Unmarshal(w.Body.Bytes(), &responseAlbums)
		require.NoError(t, err)

		assert.Equal(t, len(expectedAlbums), len(responseAlbums))
		assert.Equal(t, expectedAlbums[0].ID, responseAlbums[0].ID)
		assert.Equal(t, expectedAlbums[1].Title, responseAlbums[1].Title)

		mockRepo.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {

		c, w := setupTestContext()

		mockRepo := setupMockRepo()
		defer restoreRepo()

		mockRepo.On("GetAlbums", c.Request.Context()).Return(nil, errors.New("database error"))

		GetAllAlbums(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "error")
		assert.Equal(t, "Error getting albums", response["error"])

		mockRepo.AssertExpectations(t)
	})
}

func TestGetAlbumByID(t *testing.T) {
	t.Run("successful retrieval", func(t *testing.T) {

		c, w := setupTestContext()
		c.Params = []gin.Param{{Key: "id", Value: "1"}}

		mockRepo := setupMockRepo()
		defer restoreRepo()

		expectedAlbum := &AlbumResponse{
			ID: "1", Title: "Album One", Artist: "Artist One", Price: 9.99, ImageURL: "url1.jpg",
		}

		mockRepo.On("GetAlbum", c.Request.Context(), "1").Return(expectedAlbum, nil)

		GetAlbumByID(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseAlbum AlbumResponse
		err := json.Unmarshal(w.Body.Bytes(), &responseAlbum)
		require.NoError(t, err)

		assert.Equal(t, expectedAlbum.ID, responseAlbum.ID)
		assert.Equal(t, expectedAlbum.Title, responseAlbum.Title)
		assert.Equal(t, expectedAlbum.Artist, responseAlbum.Artist)
		assert.Equal(t, expectedAlbum.Price, responseAlbum.Price)
		assert.Equal(t, expectedAlbum.ImageURL, responseAlbum.ImageURL)

		mockRepo.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {

		c, w := setupTestContext()
		c.Params = []gin.Param{{Key: "id", Value: "999"}}

		mockRepo := setupMockRepo()
		defer restoreRepo()

		mockRepo.On("GetAlbum", c.Request.Context(), "999").Return(nil, errors.New("not found"))

		GetAlbumByID(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "error")
		assert.Equal(t, "Error getting album with ID 999", response["error"])

		mockRepo.AssertExpectations(t)
	})
}

func TestPostAlbum(t *testing.T) {
	t.Run("successful creation", func(t *testing.T) {

		c, w := setupTestContext()

		newAlbum := Album{
			Title: "New Album", Artist: "New Artist", Price: 14.99, ImageURL: "new.jpg",
		}

		jsonBytes, err := json.Marshal(newAlbum)
		require.NoError(t, err)

		c.Request = httptest.NewRequest("POST", "/albums", bytes.NewBuffer(jsonBytes))
		c.Request.Header.Set("Content-Type", "application/json")

		mockRepo := setupMockRepo()
		defer restoreRepo()

		expectedResponse := &AlbumResponse{
			ID: "3", Title: "New Album", Artist: "New Artist", Price: 14.99, ImageURL: "new.jpg",
		}

		mockRepo.On("CreateAlbum", c.Request.Context(), mock.MatchedBy(func(a Album) bool {
			return a.Title == newAlbum.Title &&
				a.Artist == newAlbum.Artist &&
				a.Price == newAlbum.Price &&
				a.ImageURL == newAlbum.ImageURL
		})).Return(expectedResponse, nil)

		PostAlbum(c)

		assert.Equal(t, http.StatusCreated, w.Code)

		var responseAlbum AlbumResponse
		err = json.Unmarshal(w.Body.Bytes(), &responseAlbum)
		require.NoError(t, err)

		assert.Equal(t, expectedResponse.ID, responseAlbum.ID)
		assert.Equal(t, expectedResponse.Title, responseAlbum.Title)
		assert.Equal(t, expectedResponse.Artist, responseAlbum.Artist)
		assert.Equal(t, expectedResponse.Price, responseAlbum.Price)
		assert.Equal(t, expectedResponse.ImageURL, responseAlbum.ImageURL)

		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid JSON", func(t *testing.T) {

		c, w := setupTestContext()

		c.Request = httptest.NewRequest("POST", "/albums", bytes.NewBuffer([]byte(`{invalid json}`)))
		c.Request.Header.Set("Content-Type", "application/json")

		mockRepo := setupMockRepo()
		defer restoreRepo()

		PostAlbum(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "error")
		assert.Equal(t, "Error with album JSON", response["error"])

		mockRepo.AssertNotCalled(t, "CreateAlbum")
	})

	t.Run("database error", func(t *testing.T) {

		c, w := setupTestContext()

		newAlbum := Album{
			Title: "Error Album", Artist: "Error Artist", Price: 14.99, ImageURL: "error.jpg",
		}

		jsonBytes, err := json.Marshal(newAlbum)
		require.NoError(t, err)

		c.Request = httptest.NewRequest("POST", "/albums", bytes.NewBuffer(jsonBytes))
		c.Request.Header.Set("Content-Type", "application/json")

		mockRepo := setupMockRepo()
		defer restoreRepo()

		mockRepo.On("CreateAlbum", c.Request.Context(), mock.MatchedBy(func(a Album) bool {
			return a.Title == newAlbum.Title &&
				a.Artist == newAlbum.Artist &&
				a.Price == newAlbum.Price &&
				a.ImageURL == newAlbum.ImageURL
		})).Return(nil, errors.New("database error"))

		PostAlbum(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]string
		err = json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "error")
		assert.Equal(t, "Error adding album", response["error"])

		mockRepo.AssertExpectations(t)
	})
}

func TestPutAlbumByID(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {

		c, w := setupTestContext()
		c.Params = []gin.Param{{Key: "id", Value: "1"}}

		updateAlbum := AlbumResponse{
			ID: "1", Title: "Updated Album", Artist: "Updated Artist", Price: 19.99, ImageURL: "updated.jpg",
		}

		jsonBytes, err := json.Marshal(updateAlbum)
		require.NoError(t, err)

		c.Request = httptest.NewRequest("PUT", "/albums/1", bytes.NewBuffer(jsonBytes))
		c.Request.Header.Set("Content-Type", "application/json")

		mockRepo := setupMockRepo()
		defer restoreRepo()

		expectedResponse := &AlbumResponse{
			ID: "1", Title: "Updated Album", Artist: "Updated Artist", Price: 19.99, ImageURL: "updated.jpg",
		}

		mockRepo.On("UpdateAlbum", c.Request.Context(), "1", updateAlbum).Return(expectedResponse, nil)

		PutAlbumByID(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseAlbum AlbumResponse
		err = json.Unmarshal(w.Body.Bytes(), &responseAlbum)
		require.NoError(t, err)

		assert.Equal(t, expectedResponse.ID, responseAlbum.ID)
		assert.Equal(t, expectedResponse.Title, responseAlbum.Title)
		assert.Equal(t, expectedResponse.Artist, responseAlbum.Artist)
		assert.Equal(t, expectedResponse.Price, responseAlbum.Price)
		assert.Equal(t, expectedResponse.ImageURL, responseAlbum.ImageURL)

		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid JSON", func(t *testing.T) {

		c, w := setupTestContext()
		c.Params = []gin.Param{{Key: "id", Value: "1"}}

		c.Request = httptest.NewRequest("PUT", "/albums/1", bytes.NewBuffer([]byte(`{invalid json}`)))
		c.Request.Header.Set("Content-Type", "application/json")

		mockRepo := setupMockRepo()
		defer restoreRepo()

		PutAlbumByID(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "error")
		assert.Equal(t, "Error with album JSON", response["error"])

		mockRepo.AssertNotCalled(t, "UpdateAlbum")
	})

	t.Run("database error", func(t *testing.T) {

		c, w := setupTestContext()
		c.Params = []gin.Param{{Key: "id", Value: "999"}}

		updateAlbum := AlbumResponse{
			ID: "999", Title: "Error Album", Artist: "Error Artist", Price: 19.99, ImageURL: "error.jpg",
		}

		jsonBytes, err := json.Marshal(updateAlbum)
		require.NoError(t, err)

		c.Request = httptest.NewRequest("PUT", "/albums/999", bytes.NewBuffer(jsonBytes))
		c.Request.Header.Set("Content-Type", "application/json")

		mockRepo := setupMockRepo()
		defer restoreRepo()

		mockRepo.On("UpdateAlbum", c.Request.Context(), "999", updateAlbum).Return(nil, errors.New("database error"))

		PutAlbumByID(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]string
		err = json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "error")
		assert.Equal(t, "Error updating album", response["error"])

		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteAlbumByID(t *testing.T) {
	t.Run("successful deletion", func(t *testing.T) {

		c, w := setupTestContext()
		c.Params = []gin.Param{{Key: "id", Value: "1"}}

		mockRepo := setupMockRepo()
		defer restoreRepo()

		expectedResponse := &AlbumResponse{
			ID: "1", Title: "Deleted Album", Artist: "Deleted Artist", Price: 9.99, ImageURL: "deleted.jpg",
		}

		mockRepo.On("DeleteAlbum", c.Request.Context(), "1").Return(expectedResponse, nil)

		DeleteAlbumByID(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseAlbum AlbumResponse
		err := json.Unmarshal(w.Body.Bytes(), &responseAlbum)
		require.NoError(t, err)

		assert.Equal(t, expectedResponse.ID, responseAlbum.ID)
		assert.Equal(t, expectedResponse.Title, responseAlbum.Title)
		assert.Equal(t, expectedResponse.Artist, responseAlbum.Artist)
		assert.Equal(t, expectedResponse.Price, responseAlbum.Price)
		assert.Equal(t, expectedResponse.ImageURL, responseAlbum.ImageURL)

		mockRepo.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {

		c, w := setupTestContext()
		c.Params = []gin.Param{{Key: "id", Value: "999"}}

		mockRepo := setupMockRepo()
		defer restoreRepo()

		mockRepo.On("DeleteAlbum", c.Request.Context(), "999").Return(nil, errors.New("database error"))

		DeleteAlbumByID(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "error")
		assert.Equal(t, "Error deleting album", response["error"])

		mockRepo.AssertExpectations(t)
	})
}
