package model

import (
)

type post struct {
	id      string
	title   string
	content string
}

func getPosts() []post {
	posts := []post{
		{
			id:      "primerPost",
			title:   "Primer Post",
			content: "Este es el primer post",
		},
		{
			id:      "segundoPost",
			title:   "Segundo Post",
			content: "Este es el segundo post",
		},
	}
	return posts
}
