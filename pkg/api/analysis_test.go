package api

import (
	"testing"

	models "github.com/CatalinCosma/UpChallenge/pkg/domain"
)

func TestAnalyzePosts(t *testing.T) {
	posts := []models.Data{
		{InstagramMedia: &models.InstagramMedia{Likes: 10, Comments: 5, TimeStamp: 2}},
		{Pin: &models.Pin{Likes: 20, Comments: 15, TimeStamp: 3}},
		{TikTokVideo: &models.TikTokVideo{Likes: 30, Comments: 25, TimeStamp: 1}},
	}

	result := AnalyzePosts(posts)

	if result.TotalPosts != len(posts) {
		t.Errorf("TotalPosts was incorrect, got: %d, want: %d.", result.TotalPosts, len(posts))
	}

	if result.MinimumTimestamp != 1 {
		t.Errorf("MinimumTimestamp was incorrect, got: %d, want: %d.", result.MinimumTimestamp, 1)
	}

	if result.MaximumTimestamp != 3 {
		t.Errorf("MaximumTimestamp was incorrect, got: %d, want: %d.", result.MaximumTimestamp, 3)
	}

	expectedAvgLikes := (10 + 20 + 30) / 3
	if result.AvgLikes != expectedAvgLikes {
		t.Errorf("AvgLikes was incorrect, got: %d, want: %d.", result.AvgLikes, expectedAvgLikes)
	}

	expectedAvgComments := (5 + 15 + 25) / 3
	if result.AvgComments != expectedAvgComments {
		t.Errorf("AvgComments was incorrect, got: %d, want: %d.", result.AvgComments, expectedAvgComments)
	}
}
