package model

type Video struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Channel     string `json:"channel"`
	Duration    string `json:"duration"`
	Views       string `json:"views"`
	PublishTime string `json:"publish_time"`
	URL         string `json:"url"`
}
