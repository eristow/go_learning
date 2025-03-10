package album

import (
	"context"
	"fmt"
	"go_learning/db"
	"os"
)

// Repository provides an interface for album data operations
type Repository interface {
	GetAlbums(ctx context.Context) ([]*AlbumResponse, error)
	GetAlbum(ctx context.Context, id string) (*AlbumResponse, error)
	CreateAlbum(ctx context.Context, album Album) (*AlbumResponse, error)
	UpdateAlbum(ctx context.Context, id string, album AlbumResponse) (*AlbumResponse, error)
	DeleteAlbum(ctx context.Context, id string) (*AlbumResponse, error)
}

// DatabaseRepo implements Repository using a database connection
type DatabaseRepo struct {
	db db.Database
}

// NewDatabaseRepo creates a database-backed repository
func NewDatabaseRepo(db db.Database) *DatabaseRepo {
	return &DatabaseRepo{db: db}
}

// GetAlbums returns all albums from the database
func (r *DatabaseRepo) GetAlbums(ctx context.Context) ([]*AlbumResponse, error) {
	query := `SELECT * FROM albums;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var albums []*AlbumResponse

	for rows.Next() {
		album := new(AlbumResponse)

		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.ImageURL)

		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}

	return albums, nil
}

// GetAlbum returns a single album by ID
func (r *DatabaseRepo) GetAlbum(ctx context.Context, id string) (*AlbumResponse, error) {
	query := `SELECT * FROM albums WHERE id = $1;`

	row := r.db.QueryRow(ctx, query, id)

	album := new(AlbumResponse)

	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.ImageURL)

	if err != nil {
		return nil, err
	}

	return album, nil
}

// CreateAlbum adds a new album to the database
func (r *DatabaseRepo) CreateAlbum(ctx context.Context, album Album) (*AlbumResponse, error) {
	query := `INSERT INTO albums (id, title, artist, price, image_url)
						VALUES (DEFAULT, $1, $2, $3, $4) RETURNING id, title, artist, price, image_url;`

	row := r.db.QueryRow(ctx, query, album.Title, album.Artist, album.Price, album.ImageURL)

	createdAlbum := new(AlbumResponse)

	err := row.Scan(&createdAlbum.ID, &createdAlbum.Title, &createdAlbum.Artist, &createdAlbum.Price, &createdAlbum.ImageURL)

	if err != nil {
		return nil, err
	}

	return createdAlbum, nil
}

// UpdateAlbum updates an existing album
func (r *DatabaseRepo) UpdateAlbum(ctx context.Context, id string, album AlbumResponse) (*AlbumResponse, error) {
	query := `UPDATE albums SET title = $1, artist = $2, price = $3, image_url = $4 WHERE id = $5 RETURNING id, title, artist, price, image_url;`

	row := r.db.QueryRow(ctx, query, album.Title, album.Artist, album.Price, album.ImageURL, id)

	updatedAlbum := new(AlbumResponse)

	err := row.Scan(&updatedAlbum.ID, &updatedAlbum.Title, &updatedAlbum.Artist, &updatedAlbum.Price, &updatedAlbum.ImageURL)

	if err != nil {
		return nil, err
	}

	return updatedAlbum, nil
}

// DeleteAlbum removes an album from the database
func (r *DatabaseRepo) DeleteAlbum(ctx context.Context, id string) (*AlbumResponse, error) {
	query := `DELETE FROM albums WHERE id = $1 RETURNING id, title, artist, price, image_url;`

	row := r.db.QueryRow(ctx, query, id)

	deletedAlbum := new(AlbumResponse)

	err := row.Scan(&deletedAlbum.ID, &deletedAlbum.Title, &deletedAlbum.Artist, &deletedAlbum.Price, &deletedAlbum.ImageURL)

	if err != nil {
		return nil, err
	}

	return deletedAlbum, nil
}
