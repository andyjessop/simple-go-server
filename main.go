package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Post struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

var (
	posts   = make(map[int]Post)
	nextID  = 1
	postsMu sync.Mutex
)

func main() {
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/posts/", postHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetPosts(w, r)
	case "POST":
		handlePostPosts(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		handleGetPost(w, r, id)
	case "DELETE":
		handleDeletePost(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetPosts(w http.ResponseWriter, r *http.Request) {
	postsMu.Lock()
	defer postsMu.Unlock()

	ps := make([]Post, 0, len(posts))
	for _, p := range posts {
		ps = append(ps, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ps)
}

func handlePostPosts(w http.ResponseWriter, r *http.Request) {
	var p Post
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	postsMu.Lock()
	defer postsMu.Unlock()

	p.ID = nextID
	nextID++
	posts[p.ID] = p

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func handleGetPost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	p, ok := posts[id]
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	_, ok := posts[id]
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	delete(posts, id)
	w.WriteHeader(http.StatusOK)
}
