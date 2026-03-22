package main

import (
	"encoding/json"
	"net/http"
)

// when the browser visits /, it just serves your index.html file back. Simple!
func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Security-Policy", "")
	http.ServeFile(w, r, "./static/index.html")
}

// when the browser hits /joke, it calls getRandomDadJoke() from jokes.go, then sends the joke back as JSON
func handleJoke(w http.ResponseWriter, r *http.Request) {
	joke, err := getRandomDadJoke()
	if err != nil {
		http.Error(w, "Failed to fetch joke", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(joke)
}

/*
What's going on:
	1. w http.ResponseWriter — this is how you write back to the browser (the response)
	2. r *http.Request — this is the incoming request from the browser.
	The * means it's a pointer — don't worry too much about this for now
	3. http.StatusInternalServerError — this is just Go's way of saying HTTP status code 500
	4. w.Header().Set("Content-Type", "application/json") —
	tells the browser "hey, what I'm sending back is JSON"

*/
