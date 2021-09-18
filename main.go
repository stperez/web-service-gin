package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Index struct {
	I string `json:"i"`
	F string `json:"f"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var index = Index{I: "Hello", F: "World"}

func main() {
	// Initialize routers with gin Default
	router := gin.Default()

	// Routes GET
	router.GET("/", Inicio)
	router.GET("/albums", getAlbums)

	// Get album by ID
	router.GET("/albums/:id", getAlbumByID)

	// Routes POST
	router.POST("/albums", postAlbums)

	// Run app
	router.Run("localhost:5050")

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Handler to return specific item
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for an album whose id value matches the parameter
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from json received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call bindjson to bind received json to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// Index view
func Inicio(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, index)
}
