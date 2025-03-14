package post

import (
)

type Post struct {
	Id      string
	Title   string
	Content string
}

func GetPosts() []Post {
	posts := []Post{
		{
			Id:      "primerPost",
			Title:   "Primer Post",
			Content: "Este es el primer Post",
		},
		{
			Id:      "segundoPost",
			Title:   "Segundo Post",
			Content: "Este es el segundo Post",
		},
	}
	return posts
}

func GetById(id string) Post {
	posts := GetPosts()
	for _, post := range posts {
		if post.Id == id {
			return post
		}
	}

	return Post{}
}
