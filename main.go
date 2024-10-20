package main

import (
    "log"
    "net/http"
    "delivery_fee_api/handlers"
)

func main() {
    http.HandleFunc("/calculate-delivery-fee", handlers.HandleDeliveryFeeRequest)
    log.Println("Starting server on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
