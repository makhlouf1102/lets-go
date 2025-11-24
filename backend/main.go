package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Problem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/problems", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "all problems",
			"data": []Problem{
				{
					ID:          "1",
					Title:       "Problem 1",
					Description: "Description 1",
					Difficulty:  "Easy",
				},
				{
					ID:          "2",
					Title:       "Problem 2",
					Description: "Description 2",
					Difficulty:  "Medium",
				},
				{
					ID:          "3",
					Title:       "Problem 3",
					Description: "Description 3",
					Difficulty:  "Hard",
				},
			},
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to run server:", err)
	}
}
