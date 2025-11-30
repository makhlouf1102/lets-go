package main

import (
	"bytes"
	"encoding/json"
	"io"
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

type RunCodeRequest struct {
	SourceCode string `json:"source_code"`
	LanguageID int    `json:"language_id"`
}

type RunCodeResponseData struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Time   string `json:"time"`
	Memory int    `json:"memory"`
	Token  string `json:"token"`
	Status struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	} `json:"status"`
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

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}

func getProblems(router *gin.Engine) *gin.Engine {
	router.GET("/problems", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "all problems",
			"data":    problems,
		})
	})
	return router
}

func getProblemById(router *gin.Engine) *gin.Engine {
	router.GET("/problems/:id", func(c *gin.Context) {
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
	return router
}

func runCode(router *gin.Engine) *gin.Engine {
	// js id is 63
	// 	{
	//   "source_code": "#include <stdio.h>\n\nint main(void) {\n  char name[10];\n  scanf(\"%s\", name);\n  printf(\"hello, %s\n\", name);\n  return 0;\n}",
	//   "language_id": 4,
	//   "stdin": "world"
	// }
	router.POST("/code/run", func(c *gin.Context) {
		var req RunCodeRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid request",
			})
			return
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to encode request",
			})
			return
		}

		// send a request to the judge
		resp, err := http.Post("http://server:2358/submissions/?base64_encoded=false&wait=true", "application/json", &buf)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to run code",
				"error":   err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to read response",
				"error":   err.Error(),
			})
			return
		}

		var response RunCodeResponseData
		if err := json.Unmarshal(respBody, &response); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to unmarshal response",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "code running successfully",
			"data":    response,
		})
	})
	return router
}
func main() {
	r := setupRouter()

	r = getProblems(r)
	r = getProblemById(r)
	r = runCode(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to run server:", err)
	}
}
