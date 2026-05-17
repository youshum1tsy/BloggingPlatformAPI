package main

import (
	"fmt"
	"log"
	"net/http"
)

func postPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create post")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get list of posts")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "get post, id: %s", id)
}

func putPost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "put post, id: %s", id)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "delete post, id: %s", id)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /posts", postPost)
	mux.HandleFunc("GET /posts", getPosts)
	mux.HandleFunc("GET /posts/{id}", getPost)
	mux.HandleFunc("PUT /posts/{id}", putPost)
	mux.HandleFunc("DELETE /posts/{id}", deletePost)
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
