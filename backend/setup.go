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
var BASE_URL string = "/api/v1"

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

	router.GET(BASE_URL+"/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	router.GET(BASE_URL+"/albums", album.GetAllAlbums)
	router.GET(BASE_URL+"/albums/:id", album.GetAlbumByID)
	router.PUT(BASE_URL+"/albums/:id", album.PutAlbumByID)
	router.POST(BASE_URL+"/albums", album.PostAlbum)
	router.DELETE(BASE_URL+"/albums/:id", album.DeleteAlbumByID)

	return router
}
