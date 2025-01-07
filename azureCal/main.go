package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/graph/armgraph"
)

func main() {
    tenantID := os.Getenv("AZURE_TENANT_ID")
    clientID := os.Getenv("AZURE_CLIENT_ID")
    clientSecret := os.Getenv("AZURE_CLIENT_SECRET")

    cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
    if err != nil {
        log.Fatalf("failed to obtain credential: %v", err)
    }

    client := armgraph.NewClient(cred, nil)

    ctx := context.Background()

    // Replace with your user's email or user ID
    userID := "user@example.com"

    events, err := client.EventsClient().List(ctx, userID, nil)
    if err != nil {
        log.Fatalf("failed to list events: %v", err)
    }

    for _, event := range events.Events {
        fmt.Printf("Event: %s\n", *event.Subject)
        fmt.Printf("Start Time: %s\n", event.Start.DateTime)
        fmt.Printf("End Time: %s\n", event.End.DateTime)
        fmt.Println("-------------------------")
    }
}
