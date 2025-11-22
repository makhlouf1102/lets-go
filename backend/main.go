package main 
import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err:= r.Run(":8080"); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}
