package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	blog := New()

	http.HandleFunc("/blog", blog.indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (b *Blog) indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/blog" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodPost {
		article := &Article{}
		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			log.Println(err)
		}

		b.SaveArticle(*article)
	}

	if r.Method == http.MethodGet {
		err := json.NewEncoder(w).Encode(b)
		if err != nil {
			log.Println(err)
		}
	}
}
