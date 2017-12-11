package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

type Redirect struct { // Our example struct, you can use "-" to ignore a field
	Url_Origin      string `csv:"url_origin"`
	Url_Destination string `csv:"url_destination"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")

	redirectsFile, err := os.OpenFile("redirects.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer redirectsFile.Close()

	redirects := []*Redirect{}

	if err := gocsv.UnmarshalFile(redirectsFile, &redirects); err != nil { // Load clients from file
		panic(err)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", func(c *gin.Context) {
		for _, redir := range redirects {
			if c.Request.Host == redir.Url_Origin {
				c.Redirect(301, redir.Url_Destination)
			}
		}
	})

	router.Run(":" + port)
}
