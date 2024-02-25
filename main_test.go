package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHandleGetPosts(t *testing.T) {
	// Setup
	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postsHandler)

	// Execute
	handler.ServeHTTP(rr, req)

	// Assert
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "[]\n" // Expect an empty JSON array
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHandlePostAndDeletePosts(t *testing.T) {
	// Create a new post
	post := &Post{
		Body: "Test post",
	}
	postBody, _ := json.Marshal(post)
	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(postBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postsHandler)

	// Execute
	handler.ServeHTTP(rr, req)

	// Assert POST
	status := rr.Code
	if status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Decode the response to get the post ID
	var newPost Post
	if err := json.NewDecoder(rr.Body).Decode(&newPost); err != nil {
		t.Fatal(err)
	}
	if newPost.Body != post.Body {
		t.Errorf("handler returned unexpected body: got %v want %v", newPost.Body, post.Body)
	}

	// Delete the post
	req, err = http.NewRequest("DELETE", "/posts/"+strconv.Itoa(newPost.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	// Execute
	handler = http.HandlerFunc(postHandler)
	handler.ServeHTTP(rr, req)

	// Assert DELETE
	status = rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code for delete: got %v want %v", status, http.StatusOK)
	}
}
