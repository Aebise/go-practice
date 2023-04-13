package models

// Post is a model that holds post data.
type Post struct {
	ID          string `json:"id"`
	ImageURL    string `json:"image_url"`
	VideoURL    string `json:"video_url"`
	Views       int64  `json:"views"`
	Like        int64  `json:"like"`
	PosterID    string `json:"poster_id"`
	CategoryID  string `json:"category_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
