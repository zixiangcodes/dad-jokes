package main

import (
	"encoding/json"
	"net/http"
)

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomDadJoke() (Joke, error) {
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		return Joke{}, err
	}

	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Joke{}, err
	}
	defer resp.Body.Close()

	var joke Joke
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		return Joke{}, err
	}

	return joke, nil

	// getRandomDadJoke() now returns the joke and an error (Joke, error) instead of printing directly — this is because the handler will be the one sending it to the browser, not this function
	// fmt is no longer needed here so you can remove it — Go will yell at you if you import something unused!
}

// Old version (prints to console directly, no error handling):
/*

func getRandomDadJoke() {
	// builds HTTP GET request to https://icanhazdadjoke.com
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		fmt.Println("ERROR: Failed to create HTTP request:", err)
		return
	}

	// tells API we want JSON object back, otherwise it returns plain text
	req.Header.Set("Accept", "application/json")

	// call the API to get a random dad joke
	fmt.Println("\nFetching a dad joke...")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: Failed to make API call:", err)
		return
	} else {
		fmt.Println("SUCCESS: The API call was successful!")
	}
	// delays the closing of the response body until the function finishes,
	// to prevent resource or memory leaks
	defer resp.Body.Close()

	// store data from API in 'Joke' struct
	var joke Joke
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		fmt.Println("ERROR: Could not read the joke:", err)
		return
	}

	// print the Dad Joke to the console for user
	fmt.Println("\n😂", joke.Joke)

	// How getRandomDadJoke() works, in-line documentation:
	//
		1. http.NewRequest — builds the HTTP GET request
		2. req.Header.Set("Accept", "application/json") — tells the API we want JSON back, otherwise it returns plain text
		3. http.Client{} — Go's built-in HTTP client, like fetch in JS or requests in Python
		4. defer resp.Body.Close() — defer means "run this when the function ends", used here to clean up the response. Very Go!
		5. json.NewDecoder(...).Decode(&joke) — unpacks the JSON into our Joke struct.
			The & passes a pointer to joke so the decoder can fill it in
		6. Every if err != nil — this is Go's error handling. No try/catch, just check and handle immediately

	}
*/
