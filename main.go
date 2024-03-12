// File: main.go

package main

import (
    "fmt"
    "net/http"

    // "mes/domain"
    "mes/delivery"
    "mes/infra"
    "mes/usecase"
)

func main() {
    // Initialize the database connection
    db := infra.GetDB()
    defer db.Close()

    // Create a DBHandler instance
    dbHandler := &delivery.DBHandler{DB: db}

    // Register HTTP handlers
    http.HandleFunc("/create", dbHandler.CreateHandler(usecase.Create))
    http.HandleFunc("/read", dbHandler.ReadHandler(usecase.Read))
    http.HandleFunc("/readall", dbHandler.ReadAllHandler(usecase.ReadAll))
    http.HandleFunc("/update", dbHandler.UpdateHandler(usecase.Update))
    http.HandleFunc("/delete", dbHandler.DeleteHandler(usecase.Delete))

    // Start the HTTP server
    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}
