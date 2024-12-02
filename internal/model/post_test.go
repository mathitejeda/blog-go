package model

import (
	"testing"
)

func TestGetPosts(t *testing.T) {
	posts := getPosts()
	for _, post := range posts {
		t.Log(post.id)
		t.Log(post.title)
		t.Log(post.content)
	}
	if len(posts) != 2 {
		t.Errorf("Expected 2 posts, got %d", len(posts))
	}
}
