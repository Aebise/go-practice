package models

// PostNotification holds data about an event to inform the user related with a post
type PostNotification struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	PerformedBy string `json:"performed_by"`
	PostID      string `json:"post_id"`
	Message     string `json:"message"`
}

// Need to have account related notifications.
