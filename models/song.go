package models

type Song struct {
	Id      int    `json:"id"`
	Theme   string `json:"theme"`
	Artiste string `json:"artiste"`
	Song    string `json:"song"`
	Image   string `json:"image"`
	Radio   string `json:"radio"`
	Title   string `json:"title"`
}
