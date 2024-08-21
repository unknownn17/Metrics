package models

type Create struct {
	Title          string `json:"title"`
	Author         string `json:"author"`
	Published_Year int `json:"published_year"`
}

type Update struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Published_Year int `json:"published_year"`
}
