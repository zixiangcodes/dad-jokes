package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/joke", handleJoke)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default to 8080 if PORT environment variable is not set
		openBrowser("http://localhost:" + port)
	}

	fmt.Println("Server starting on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("ERROR: Could not start server:", err)
	}
}

func openBrowser(url string) {
	switch runtime.GOOS {
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	case "linux":
		exec.Command("xdg-open", url).Start()
	}
}

/*
What's going on:

1. http.HandleFunc("/", handleIndex) — registers the route. When browser visits /, call handleIndex

2. http.HandleFunc("/joke", handleJoke) — same but for /joke, calls handleJoke

3. os.Getenv("PORT") — reads the PORT environment variable. On Railway/Render, this is automatically set for you
	if port == "" — if there's no PORT env variable (i.e. you're running locally), default to 8080
	Note: This is actually the industry standard pattern for this — you'll see it in almost every production Go app.

4. http.ListenAndServe(":"+port, nil) — starts the server on the specified port and listens for incoming requests

5. nil here means "use the default router" — no need for anything custom

*/

/*
func main() {
	startAnnouncement()
	welcomeAnnouncement()
	name := askUserName()
	getRandomDadJoke()
	askForAnotherJoke(name)
}

// list of functions to be called in main()
// Note: CLI version

func startAnnouncement() {
	fmt.Println("\nStarting the dad joke generator app...")
}

func welcomeAnnouncement() {
	fmt.Println("Welcome to the dad joke generator!")
}

func askUserName() string {
	var name string
	fmt.Print("\nEnter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
	return name
}

func askForAnotherJoke(name string) {
	var response string
	fmt.Print("\nWould you like another dad joke? (y/n): ")
	fmt.Scanln(&response)
	if response == "y" || response == "Y" {
		getRandomDadJoke()
		askForAnotherJoke(name)
	} else {
		quitAnnouncement(name)
	}
}

func quitAnnouncement(name string) {
	fmt.Println("Thank you ", name, " for using the dad joke generator...")
	fmt.Println("Quitting the dad joke generator...")
}
*/
