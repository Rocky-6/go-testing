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

func TestIndexHandler(t *testing.T) {
	blog := New()

	http.HandleFunc("/blog", blog.indexHandler)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/blog", strings.NewReader(`{"title": "xxx", "body": "xxxx"}`))
	blog.indexHandler(w, req)

	assert.Equal(t, w.Code, 200)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/blog", nil)
	blog.indexHandler(w, req)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), "{\"articles\":[{\"title\":\"xxx\",\"body\":\"xxxx\"}]}\n")
}
