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
	dbHandler := usecase.NewDBHandler(db)
    // Register HTTP handlers
	http.HandleFunc("/create", delivery.CreateHandler(dbHandler.Create))
    http.HandleFunc("/read", delivery.ReadHandler(dbHandler.Read))
    http.HandleFunc("/readall", delivery.ReadAllHandler(dbHandler.ReadAll))
    http.HandleFunc("/update", delivery.UpdateHandler(dbHandler.Update))
    http.HandleFunc("/delete", delivery.DeleteHandler(dbHandler.Delete))

    // Start the HTTP server
    fmt.Println("Server listening on port 8082...")
    http.ListenAndServe(":8082", nil)
}
