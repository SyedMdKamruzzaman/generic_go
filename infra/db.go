// File: infra/db.go

package infra

import (
    "fmt"
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
    var err error
    // Configure PostgreSQL connection
    db, err = gorm.Open("postgres", "user=postgres password=123456 dbname=mes sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to the database")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
    return db
}
