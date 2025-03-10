package main

import (
	"context"
	"fmt"
	"go_learning/album"
	"go_learning/db"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Router interface {
	Run(addr ...string) error
}

type PGXSetupFunc func() db.Database
type RouterSetupFunc func() Router

var SetupPGX PGXSetupFunc = setupPGX
var SetupRouter RouterSetupFunc = setupRouter

func setupPGX() db.Database {
	dbURL := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), dbURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func setupRouter() Router {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	// TODO: change CORS rules to only allow docker containers
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	router.GET("/albums", album.GetAllAlbums)
	router.GET("/albums/:id", album.GetAlbumByID)
	router.PUT("/albums/:id", album.PutAlbumByID)
	router.POST("/albums", album.PostAlbum)
	router.DELETE("/albums/:id", album.DeleteAlbumByID)

	return router
}
