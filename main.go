package main

import (
	"github.com/drone/gin-location"
	"github.com/gin-gonic/gin"
	"log"
	//	"net/http"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
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

	// configure to automatically detect scheme and host
	// - use http when default scheme cannot be determined
	// - use localhost:8080 when default host cannot be determined
	router.Use(location.Default())

	redirectsFile, err := os.OpenFile("redirects.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer redirectsFile.Close()

	redirects := []*Redirect{}

	if err := gocsv.UnmarshalFile(redirectsFile, &redirects); err != nil { // Load clients from file
		panic(err)
	}

	router.GET("/", func(c *gin.Context) {
		url := location.Get(c)
		log.Printf(url.Host)

		for _, redir := range redirects {
			fmt.Println("Hello", redir.Url_Destination)
		}
		// url.Scheme
		// url.Host
		// url.Path
	})

	router.Run(":" + port)
}
