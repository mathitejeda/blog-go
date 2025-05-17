package post

import (
	"fmt"
	"github.com/niklasfasching/go-org/org"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Post struct {
	Id      string
	Title   string
	Content string
	Date    time.Time
	Summary string
}

const postBaseDir = "./posts"
const dateLayout = "2006-01-02"

func GetPosts() []Post {
	dir := postBaseDir
	config := org.New()
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var posts []Post

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".org") {
			continue
		}

		id := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		path := filepath.Join(dir, entry.Name())

		data, err := os.ReadFile(path)
		orgData := config.Parse(strings.NewReader(string(data)), "")

		if err != nil {
			fmt.Printf("Error leyendo %s: %v\n", path, err)
			continue
		}

		title := orgData.Get("TITLE")
		dateS := orgData.Get("DATE")
		summary := orgData.Get("SUMMARY")

		date, err := time.Parse(dateLayout, dateS)

		posts = append(posts, Post{
			Id:      id,
			Title:   title,
			Content: summary,
			Date:    date,
			Summary: summary,
		})
	}

	return posts
}

func GetAsHTML(file string) (string, error) {
	filepath := fmt.Sprintf("%s/%s.org", postBaseDir, file)

	content, err := os.ReadFile(filepath)

	if err != nil {
		return "", err
	}

	html, err := org.New().Parse(strings.NewReader(string(content)), "").Write(org.NewHTMLWriter())

	return html, nil
}

func (p Post)DateAsString() string {
	return p.Date.Format(dateLayout)
}
