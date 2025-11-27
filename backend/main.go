package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Problem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
}

var problems = []Problem{
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
			"data":    problems,
		})
	})

	r.GET("/problems/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid id",
			})
			return
		}

		if id < 0 || id >= len(problems) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "problem not found",
			})
			return
		}

		problem := problems[id]
		c.JSON(http.StatusOK, gin.H{
			"message": "problem",
			"data":    problem,
		})
	})

	r.POST("/code/run", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "code run",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to run server:", err)
	}
}
