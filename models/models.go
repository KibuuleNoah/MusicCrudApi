package models

type Song struct {
	ID       int    `json:"id" form:"id" uri:"id" binding:"required"`
	Title    string `json:"title" form:"title" uri:"title"`
	Artist   string `json:"artist" form:"artist" uri:"artist"`
	Genre    string `json:"genre" form:"genre" uri:"genre"`
	Released int    `json:"released" form:"released" uri:"released"`
}
