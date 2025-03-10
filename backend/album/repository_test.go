package album

import (
	"context"
	"errors"
	"go_learning/test_util"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAlbums(t *testing.T) {
	tests := []struct {
		name       string
		mockData   [][]interface{}
		mockError  error
		wantErr    bool
		wantAlbums []*AlbumResponse
	}{
		{
			name: "successful query",
			mockData: [][]interface{}{
				{"1", "Album One", "Artist One", 9.99, "url1.jpg"},
				{"2", "Album Two", "Artist Two", 12.99, "url2.jpg"},
			},
			mockError: nil,
			wantErr:   false,
			wantAlbums: []*AlbumResponse{
				{ID: "1", Title: "Album One", Artist: "Artist One", Price: 9.99, ImageURL: "url1.jpg"},
				{ID: "2", Title: "Album Two", Artist: "Artist Two", Price: 12.99, ImageURL: "url2.jpg"},
			},
		},
		{
			name:       "database error",
			mockData:   nil,
			mockError:  errors.New("database error"),
			wantErr:    true,
			wantAlbums: nil,
		},
		{
			name: "scan error",
			mockData: [][]interface{}{
				{"1", "Album One", "Artist One", 9.99, "url1.jpg"},
			},
			mockError:  errors.New("scan error"),
			wantErr:    true,
			wantAlbums: nil,
		},
		{
			name:       "empty result",
			mockData:   [][]interface{}{},
			mockError:  nil,
			wantErr:    false,
			wantAlbums: []*AlbumResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := new(test_util.MockDBConn)
			repo := NewDatabaseRepo(mockDB)
			ctx := context.Background()

			if tt.name == "database error" {
				mockDB.On("Query", ctx, "SELECT * FROM albums;").Return(nil, tt.mockError)

				albums, err := repo.GetAlbums(ctx)

				assert.Error(t, err)
				assert.Equal(t, tt.mockError, err)
				assert.Nil(t, albums)

				mockDB.AssertExpectations(t)
				return
			}

			mockRows := &test_util.MockRows{
				Data: tt.mockData,
			}

			if tt.name == "empty result" {
				mockRows.On("Next").Return(false).Once()
				mockRows.On("Close").Return()
			} else if tt.name == "scan error" {
				mockRows.On("Next").Return(true).Once()
				mockRows.On("Scan", mock.Anything).Return(tt.mockError).Once()
				mockRows.On("Close").Return()
			} else {
				for i := 0; i < len(tt.mockData); i++ {
					mockRows.On("Next").Return(true).Once()
				}
				mockRows.On("Next").Return(false).Once()

				mockRows.On("Scan", mock.Anything).Return(nil).Times(len(tt.mockData))
				mockRows.On("Close").Return()
			}

			mockDB.On("Query", ctx, "SELECT * FROM albums;").Return(mockRows, nil)

			albums, err := repo.GetAlbums(ctx)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, albums)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tt.wantAlbums), len(albums))

				if len(tt.wantAlbums) > 0 {
					for i, album := range albums {
						assert.Equal(t, tt.wantAlbums[i].ID, album.ID, "ID mismatch at index %d", i)
						assert.Equal(t, tt.wantAlbums[i].Title, album.Title, "Title mismatch at index %d", i)
						assert.Equal(t, tt.wantAlbums[i].Artist, album.Artist, "Artist mismatch at index %d", i)
						assert.Equal(t, tt.wantAlbums[i].Price, album.Price, "Price mismatch at index %d", i)
						assert.Equal(t, tt.wantAlbums[i].ImageURL, album.ImageURL, "ImageURL mismatch at index %d", i)
					}
				}
			}

			mockDB.AssertExpectations(t)
			mockRows.AssertExpectations(t)
		})
	}
}

func TestGetAlbum(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		mockData []interface{}
		mockErr  error
		wantErr  bool
		expected *AlbumResponse
	}{
		{
			name:     "successful query",
			id:       "1",
			mockData: []interface{}{"1", "Album One", "Artist One", 9.99, "url1.jpg"},
			mockErr:  nil,
			wantErr:  false,
			expected: &AlbumResponse{ID: "1", Title: "Album One", Artist: "Artist One", Price: 9.99, ImageURL: "url1.jpg"},
		},
		{
			name:     "not found",
			id:       "999",
			mockData: nil,
			mockErr:  errors.New("not found"),
			wantErr:  true,
			expected: nil,
		},
		{
			name:     "database error",
			id:       "1",
			mockData: nil,
			mockErr:  errors.New("database error"),
			wantErr:  true,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := new(test_util.MockDBConn)
			repo := NewDatabaseRepo(mockDB)
			ctx := context.Background()
			mockRow := &test_util.MockRow{
				Data: tt.mockData,
			}

			mockRow.On("Scan", mock.Anything).Return(tt.mockErr)
			mockDB.On("QueryRow", ctx,
				test_util.SQLMatcher("SELECT * FROM albums WHERE id = $1;"),
				tt.id).Return(mockRow)

			album, err := repo.GetAlbum(ctx, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, album)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, album)
				assert.Equal(t, tt.expected.ID, album.ID)
				assert.Equal(t, tt.expected.Title, album.Title)
				assert.Equal(t, tt.expected.Artist, album.Artist)
				assert.Equal(t, tt.expected.Price, album.Price)
				assert.Equal(t, tt.expected.ImageURL, album.ImageURL)
			}

			mockDB.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}

func TestCreateAlbum(t *testing.T) {
	tests := []struct {
		name     string
		album    Album
		mockData []interface{}
		mockErr  error
		wantErr  bool
		expected *AlbumResponse
	}{
		{
			name: "successful creation",
			album: Album{
				Title:    "New Album",
				Artist:   "New Artist",
				Price:    14.99,
				ImageURL: "new.jpg",
			},
			mockData: []interface{}{"3", "New Album", "New Artist", 14.99, "new.jpg"},
			mockErr:  nil,
			wantErr:  false,
			expected: &AlbumResponse{ID: "3", Title: "New Album", Artist: "New Artist", Price: 14.99, ImageURL: "new.jpg"},
		},
		{
			name: "database error",
			album: Album{
				Title:    "Error Album",
				Artist:   "Error Artist",
				Price:    14.99,
				ImageURL: "error.jpg",
			},
			mockData: nil,
			mockErr:  errors.New("database error"),
			wantErr:  true,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := new(test_util.MockDBConn)
			repo := NewDatabaseRepo(mockDB)
			ctx := context.Background()
			mockRow := &test_util.MockRow{
				Data: tt.mockData,
			}

			sql := "INSERT INTO albums (id, title, artist, price, image_url) VALUES (DEFAULT, $1, $2, $3, $4) RETURNING id, title, artist, price, image_url;"
			mockRow.On("Scan", mock.Anything).Return(tt.mockErr)
			mockDB.On("QueryRow", ctx,
				test_util.SQLMatcher(sql),
				tt.album.Title, tt.album.Artist, tt.album.Price, tt.album.ImageURL).Return(mockRow)

			album, err := repo.CreateAlbum(ctx, tt.album)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, album)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, album)
				assert.Equal(t, tt.expected.ID, album.ID)
				assert.Equal(t, tt.expected.Title, album.Title)
				assert.Equal(t, tt.expected.Artist, album.Artist)
				assert.Equal(t, tt.expected.Price, album.Price)
				assert.Equal(t, tt.expected.ImageURL, album.ImageURL)
			}

			mockDB.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}

func TestUpdateAlbum(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		album    AlbumResponse
		mockData []interface{}
		mockErr  error
		wantErr  bool
		expected *AlbumResponse
	}{
		{
			name: "successful update",
			id:   "1",
			album: AlbumResponse{
				ID:       "1",
				Title:    "Updated Album",
				Artist:   "Updated Artist",
				Price:    19.99,
				ImageURL: "updated.jpg",
			},
			mockData: []interface{}{"1", "Updated Album", "Updated Artist", 19.99, "updated.jpg"},
			mockErr:  nil,
			wantErr:  false,
			expected: &AlbumResponse{ID: "1", Title: "Updated Album", Artist: "Updated Artist", Price: 19.99, ImageURL: "updated.jpg"},
		},
		{
			name: "not found",
			id:   "999",
			album: AlbumResponse{
				ID:       "999",
				Title:    "Missing Album",
				Artist:   "Missing Artist",
				Price:    19.99,
				ImageURL: "missing.jpg",
			},
			mockData: nil,
			mockErr:  errors.New("not found"),
			wantErr:  true,
			expected: nil,
		},
		{
			name: "database error",
			id:   "1",
			album: AlbumResponse{
				ID:       "1",
				Title:    "Error Album",
				Artist:   "Error Artist",
				Price:    19.99,
				ImageURL: "error.jpg",
			},
			mockData: nil,
			mockErr:  errors.New("database error"),
			wantErr:  true,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := new(test_util.MockDBConn)
			repo := NewDatabaseRepo(mockDB)
			ctx := context.Background()
			mockRow := &test_util.MockRow{
				Data: tt.mockData,
			}

			sql := "UPDATE albums SET title = $1, artist = $2, price = $3, image_url = $4 WHERE id = $5 RETURNING id, title, artist, price, image_url;"
			mockRow.On("Scan", mock.Anything).Return(tt.mockErr)
			mockDB.On("QueryRow", ctx,
				test_util.SQLMatcher(sql),
				tt.album.Title, tt.album.Artist, tt.album.Price, tt.album.ImageURL, tt.id).Return(mockRow)

			album, err := repo.UpdateAlbum(ctx, tt.id, tt.album)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, album)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, album)
				assert.Equal(t, tt.expected.ID, album.ID)
				assert.Equal(t, tt.expected.Title, album.Title)
				assert.Equal(t, tt.expected.Artist, album.Artist)
				assert.Equal(t, tt.expected.Price, album.Price)
				assert.Equal(t, tt.expected.ImageURL, album.ImageURL)
			}

			mockDB.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}

func TestDeleteAlbum(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		mockData []interface{}
		mockErr  error
		wantErr  bool
		expected *AlbumResponse
	}{
		{
			name:     "successful deletion",
			id:       "1",
			mockData: []interface{}{"1", "Deleted Album", "Deleted Artist", 9.99, "deleted.jpg"},
			mockErr:  nil,
			wantErr:  false,
			expected: &AlbumResponse{ID: "1", Title: "Deleted Album", Artist: "Deleted Artist", Price: 9.99, ImageURL: "deleted.jpg"},
		},
		{
			name:     "not found",
			id:       "999",
			mockData: nil,
			mockErr:  errors.New("not found"),
			wantErr:  true,
			expected: nil,
		},
		{
			name:     "database error",
			id:       "1",
			mockData: nil,
			mockErr:  errors.New("database error"),
			wantErr:  true,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := new(test_util.MockDBConn)
			repo := NewDatabaseRepo(mockDB)
			ctx := context.Background()
			mockRow := &test_util.MockRow{
				Data: tt.mockData,
			}

			sql := "DELETE FROM albums WHERE id = $1 RETURNING id, title, artist, price, image_url;"
			mockRow.On("Scan", mock.Anything).Return(tt.mockErr)
			mockDB.On("QueryRow", ctx, test_util.SQLMatcher(sql), tt.id).Return(mockRow)

			album, err := repo.DeleteAlbum(ctx, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, album)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, album)
				assert.Equal(t, tt.expected.ID, album.ID)
				assert.Equal(t, tt.expected.Title, album.Title)
				assert.Equal(t, tt.expected.Artist, album.Artist)
				assert.Equal(t, tt.expected.Price, album.Price)
				assert.Equal(t, tt.expected.ImageURL, album.ImageURL)
			}

			mockDB.AssertExpectations(t)
			mockRow.AssertExpectations(t)
		})
	}
}
