package calapi

import(
	"fmt"
	"net/http"
    "github.com/gin-gonic/gin"
)

func add(c *gin.Context) {
    // Call BindJSON to bind the received JSON to
    if err := c.BindJSON(&add); err != nil {
        return 
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
