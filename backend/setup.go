package main

import (
	"context"
	"fmt"
	"go_learning/album"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupPGX() *pgx.Conn {
	dbURL := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), dbURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v/n", err)
		os.Exit(1)
	}

	return conn
}

func SetupRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	// TODO: change to only allow docker containers
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	router.GET("/albums", album.GetAlbums)
	router.GET("/albums/:id", album.GetAlbumByID)
	router.POST("/albums", album.PostAlbums)
	router.DELETE("/albums/:id", album.DeleteAlbumByID)

	return router
}
