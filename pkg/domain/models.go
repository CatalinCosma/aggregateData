package models

type InstagramMedia struct {
	ID        int `json:"id"`
	Likes     int `json:"likes"`
	Comments  int `json:"comments"`
	TimeStamp int `json:"timestamp"`
}

type Article struct {
	ID        int `json:"id"`
	Likes     int `json:"likes"`
	Comments  int `json:"comments"`
	TimeStamp int `json:"timestamp"`
}

type FacebookStatus struct {
	ID        int `json:"id"`
	Likes     int `json:"likes"`
	Comments  int `json:"comments"`
	TimeStamp int `json:"timestamp"`
}

type YoutubeVideo struct {
	ID        int `json:"id"`
	Likes     int `json:"likes"`
	Comments  int `json:"comments"`
	TimeStamp int `json:"timestamp"`
}

type TikTokVideo struct {
	ID        int `json:"id"`
	Likes     int `json:"likes"`
	Comments  int `json:"comments"`
	TimeStamp int `json:"timestamp"`
}

type Tweet struct {
	ID        int `json:"id"`
	Likes     int `json:"likes"`
	Comments  int `json:"comments"`
	TimeStamp int `json:"timestamp"`
}

type Pin struct {
	ID        int `json:"id"`
	Likes     int `json:"likes"`
	Comments  int `json:"comments"`
	TimeStamp int `json:"timestamp"`
}

type Data struct {
	InstagramMedia *InstagramMedia `json:"instagram_media,omitempty"`
	Pin            *Pin            `json:"pin,omitempty"`
	YoutubeVideo   *YoutubeVideo   `json:"youtube_video,omitempty"`
	TikTokVideo    *TikTokVideo    `json:"tiktok_video,omitempty"`
	Tweet          *Tweet          `json:"tweet,omitempty"`
	Article        *Article        `json:"article,omitempty"`
	FacebookStatus *FacebookStatus `json:"facebook_status,omitempty"`
}

type AnalysisResult struct {
	TotalPosts       int `json:"total_posts"`
	MinimumTimestamp int `json:"minimum_timestamp"`
	MaximumTimestamp int `json:"maximum_timestamp"`
	AvgLikes         int `json:"avg_likes,omitempty"`
	AvgComments      int `json:"avg_comments,omitempty"`
}

func (data *Data) GetLikes() int {
	if data.InstagramMedia != nil {
		return data.InstagramMedia.Likes
	}
	if data.Pin != nil {
		return data.Pin.Likes
	}
	if data.YoutubeVideo != nil {
		return data.YoutubeVideo.Likes
	}
	if data.TikTokVideo != nil {
		return data.TikTokVideo.Likes
	}
	if data.Tweet != nil {
		return data.Tweet.Likes
	}
	if data.Article != nil {
		return data.Article.Likes
	}
	if data.FacebookStatus != nil {
		return data.FacebookStatus.Likes
	}
	return 0
}

func (data *Data) GetComments() int {
	if data.InstagramMedia != nil {
		return data.InstagramMedia.Comments
	}
	if data.Pin != nil {
		return data.Pin.Comments
	}
	if data.YoutubeVideo != nil {
		return data.YoutubeVideo.Comments
	}
	if data.TikTokVideo != nil {
		return data.TikTokVideo.Comments
	}
	if data.Tweet != nil {
		return data.Tweet.Comments
	}
	if data.Article != nil {
		return data.Article.Comments
	}
	if data.FacebookStatus != nil {
		return data.FacebookStatus.Comments
	}
	return 0
}

func (data *Data) GetTimestamp() int {
	if data.InstagramMedia != nil {
		return data.InstagramMedia.TimeStamp
	}
	if data.Pin != nil {
		return data.Pin.TimeStamp
	}
	if data.YoutubeVideo != nil {
		return data.YoutubeVideo.TimeStamp
	}
	if data.TikTokVideo != nil {
		return data.TikTokVideo.TimeStamp
	}
	if data.Tweet != nil {
		return data.Tweet.TimeStamp
	}
	if data.Article != nil {
		return data.Article.TimeStamp
	}
	if data.FacebookStatus != nil {
		return data.FacebookStatus.TimeStamp
	}
	return 0
}
