package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveArticle(t *testing.T) {

	blog := New()

	blog.SaveArticle(Article{"My title", "My Post Body"})

	if blog.Articles[0].Title != "My title" {
		t.Errorf("Item was not added")
	}
}

func TestFetchAllArticles(t *testing.T) {

	blog := New()

	blog.SaveArticle(Article{"My title", "My Post Body"})

	articles := blog.FetchAll()

	if len(articles) == 0 {
		t.Errorf("Fetch All fails")
	}
}

func TestPostMethod(t *testing.T) {
	blog := New()

	http.HandleFunc("/blog", blog.indexHandler)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/blog", strings.NewReader(`{"title": "xxx", "body": "xxx"}`))
	blog.indexHandler(w, req)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, 200)
}
