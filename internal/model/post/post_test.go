package post

import (
	"testing"
)

func TestGetPosts(t *testing.T) {
	posts := GetPosts()
	for _, post := range posts {
		t.Log(post.Id)
		t.Log(post.Title)
		t.Log(post.Content)
	}
	if len(posts) != 2 {
		t.Errorf("Expected 2 posts, got %d", len(posts))
	}
}

func TestGetHTML(t *testing.T) {
	content, err := GetAsHTML("post")

	if err != nil {
		t.Errorf("expected post err: %s", err)
	}

	t.Log(content)
}
