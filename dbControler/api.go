package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "gin"
)

type Message struct {
    Text string `json:"text"`
}
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
// helloHandler is the HTTP handler for the /hello GET endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n %s", r.URL.Path)
    return
}
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
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

func deleteMessageHandler(w http.ResponseWriter, r *http.Request) {
    // Only accept DELETE requests
    if r.Method != http.MethodDelete {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // TODO: Delete the message from your database

    // Respond to the request
    fmt.Fprintf(w, "Message deleted")
}


func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/messages", postMessageHandler)
    http.HandleFunc("/deletemessages", deleteMessageHandler)
    http.HandleFunc("/album", getAlbums)
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
