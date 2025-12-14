package main

import "github.com/Yer01/internal/model"

type templateData struct {
	Title      string
	Year       int
	Blog       *model.Blog
	Blogs      []model.Blog
	TotalPosts int
}
