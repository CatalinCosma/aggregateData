package main

import (
	"net/http"
	"time"

	"github.com/CatalinCosma/UpChallenge/pkg/api"
	models "github.com/CatalinCosma/UpChallenge/pkg/domain"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/analysis", func(c *gin.Context) {
		durationStr := c.Query("duration")
		dimension := c.Query("dimension")

		if durationStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Duration query parameter is required"})
			return
		}

		duration, err := time.ParseDuration(durationStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duration parameter"})
			return
		}

		if dimension == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dimension query parameter is required"})
			return
		}

		postsChan := make(chan []models.Data)
		errChan := make(chan error)

		go func() {
			posts, err := api.FetchPostsData("https://stream.upfluence.co/stream", duration)
			if err != nil {
				errChan <- err
				return
			}
			postsChan <- posts
		}()

		select {
		case posts := <-postsChan:
			analysisResults := api.AnalyzePosts(posts)
			var resp models.AnalysisResult

			if dimension == "comments" {
				resp = models.AnalysisResult{
					TotalPosts:       analysisResults.TotalPosts,
					MinimumTimestamp: analysisResults.MinimumTimestamp,
					MaximumTimestamp: analysisResults.MaximumTimestamp,
					AvgComments:      analysisResults.AvgComments,
				}
			}

			if dimension == "likes" {
				resp = models.AnalysisResult{
					TotalPosts:       analysisResults.TotalPosts,
					MinimumTimestamp: analysisResults.MinimumTimestamp,
					MaximumTimestamp: analysisResults.MaximumTimestamp,
					AvgLikes:         analysisResults.AvgLikes,
				}
			}

			c.JSON(http.StatusOK, resp)
		case err := <-errChan:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	r.Run() // starting server on default port: 8080
}
