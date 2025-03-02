package album

import (
	"context"
	"fmt"
	"go_learning/db"
	"os"
)

func DbGetAlbums() ([]*AlbumResponse, error) {
	query := `SELECT * FROM albums;`

	rows, err := db.DBConn.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}
	defer rows.Close()

	var albums []*AlbumResponse

	if rows != nil {
		for rows.Next() {
			album := new(AlbumResponse)

			err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.ImageURL)

			if err != nil {
				return nil, err
			}

			albums = append(albums, album)
		}
	}

	return albums, nil
}

func DbGetAlbum(id string) (*AlbumResponse, error) {
	query := `SELECT * FROM albums WHERE id = $1;`

	row := db.DBConn.QueryRow(context.Background(), query, id)

	album := new(AlbumResponse)

	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.ImageURL)

	if err != nil {
		return nil, err
	}

	return album, nil
}

func DbCreateAlbum(album Album) (*AlbumResponse, error) {
	query := `INSERT INTO albums (id, title, artist, price, image_url)
						VALUES (DEFAULT, $1, $2, $3, $4) RETURNING id, title, artist, price, image_url;`

	row := db.DBConn.QueryRow(context.Background(), query, album.Title, album.Artist, album.Price, album.ImageURL)

	newAlbum := new(AlbumResponse)

	err := row.Scan(&newAlbum.ID, &newAlbum.Title, &newAlbum.Artist, &newAlbum.Price, &newAlbum.ImageURL)

	if err != nil {
		return nil, err
	}

	return newAlbum, nil
}

func DbDeleteAlbum(id string) (*AlbumResponse, error) {
	query := `DELETE FROM albums WHERE id = $1 RETURNING id, title, artist, price, image_url;`

	row := db.DBConn.QueryRow(context.Background(), query, id)

	album := new(AlbumResponse)

	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.ImageURL)

	if err != nil {
		return nil, err
	}

	return album, nil
}
