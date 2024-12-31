package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

type Resource struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Type        string `json:"type"`
    Description string `json:"description"`
}

type Resources struct {
    Resources []Resource `json:"resources"`
}

func readResourcesFromFile(filename string) (Resources, error) {
    file, err := os.Open(filename)
    if err != nil {
        return Resources{}, err
    }
    defer file.Close()

    byteValue, _ := ioutil.ReadAll(file)

    var resources Resources
    json.Unmarshal(byteValue, &resources)

    return resources, nil
}

func printResources(resources Resources) {
    for _, resource := range resources.Resources {
        fmt.Printf("ID: %d, Name: %s, Type: %s, Description: %s\n",
            resource.ID, resource.Name, resource.Type, resource.Description)
    }
}

func main() {
    filename := "resources.json"
    resources, err := readResourcesFromFile(filename)
    if err != nil {
        log.Fatalf("Failed to read resources from file: %v", err)
    }

    printResources(resources)
}
