package api

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"time"

	models "github.com/CatalinCosma/UpChallenge/pkg/domain"
)

func AnalyzePosts(posts []models.Data) models.AnalysisResult {
	totalPosts := len(posts)
	if totalPosts == 0 {
		return models.AnalysisResult{}
	}

	totalLikes := 0
	totalComments := 0
	minTimestamp := posts[0].GetTimestamp()
	maxTimestamp := minTimestamp

	for _, post := range posts {
		totalLikes += post.GetLikes()
		totalComments += post.GetComments()
		timestamp := post.GetTimestamp()
		if timestamp < minTimestamp {
			minTimestamp = timestamp
		}
		if timestamp > maxTimestamp {
			maxTimestamp = timestamp
		}
	}

	return models.AnalysisResult{
		TotalPosts:       totalPosts,
		MinimumTimestamp: minTimestamp,
		MaximumTimestamp: maxTimestamp,
		AvgLikes:         totalLikes / totalPosts,
		AvgComments:      totalComments / totalPosts,
	}
}

func FetchPostsData(url string, duration time.Duration) ([]models.Data, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []models.Data
	reader := bufio.NewReader(resp.Body)
	done := make(chan bool)
	go func() {
		time.Sleep(duration)
		done <- true
	}()

	for {
		select {
		case <-done:
			return posts, nil
		default:
			line, err := reader.ReadBytes('\n')
			if err != nil && err != io.EOF {
				return nil, err
			}
			if err == io.EOF {
				break
			}

			if len(line) > 6 { // assuming a minimal json object is "{}\n"
				var data models.Data
				// Parse JSON and append to posts if it is valid
				if err := json.Unmarshal(line[6:], &data); err == nil {
					posts = append(posts, data)
				}
			}
		}
	}
}
