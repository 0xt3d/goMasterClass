package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Message struct {
    Text string `json:"text"`
}

// helloHandler is the HTTP handler for the /hello GET endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

// postMessageHandler is the HTTP handler for the /messages POST endpoint
func postMessageHandler(w http.ResponseWriter, r *http.Request) {
    // Only accept POST requests
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Read the request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()

    // Parse the request body as JSON
    var msg Message
    if err := json.Unmarshal(body, &msg); err != nil {
        http.Error(w, "Error parsing request body as JSON", http.StatusBadRequest)
        return
    }

    // Do something with the message (e.g., save it to a database)
    fmt.Printf("Received message: %s\n", msg.Text)

    // Respond to the request
    fmt.Fprintf(w, "Message received")
}

func main() {
    // Tell the HTTP server to use helloHandler for the /hello GET endpoint
    http.HandleFunc("/hello", helloHandler)

    // Tell the HTTP server to use postMessageHandler for the /messages POST endpoint
    http.HandleFunc("/messages", postMessageHandler)

    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
