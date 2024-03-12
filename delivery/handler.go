// File: delivery/handler.go

package delivery

import (
    "encoding/json"
    "net/http"

    "github.com/jinzhu/gorm"
    "mes/domain"
)

// DBHandler embeds *gorm.DB to avoid passing it to every handler
type DBHandler struct {
    *gorm.DB
}

// CreateHandler handles the creation of a model
func (db *DBHandler) CreateHandler(createFunc func(*gorm.DB, *domain.Model) error) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var model domain.Model
        decoder := json.NewDecoder(r.Body)
        if err := decoder.Decode(&model); err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }
        defer r.Body.Close()

        if err := createFunc(db.DB, &model); err != nil {
            http.Error(w, "Failed to create model", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(&model)
    }
}

// ReadHandler handles the retrieval of a model by ID
func (db *DBHandler) ReadHandler(readFunc func(*gorm.DB, *domain.Model, string) error) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Query().Get("id")
        var model domain.Model
        if err := readFunc(db.DB, &model, id); err != nil {
            http.Error(w, "Model not found", http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(&model)
    }
}

// ReadAllHandler handles the retrieval of all models
func (db *DBHandler) ReadAllHandler(readAllFunc func(*gorm.DB, *[]domain.Model) error) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var models []domain.Model
        if err := readAllFunc(db.DB, &models); err != nil {
            http.Error(w, "Failed to fetch models", http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(&models)
    }
}

// UpdateHandler handles the update of a model
func (db *DBHandler) UpdateHandler(updateFunc func(*gorm.DB, *domain.Model) error) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var model domain.Model
        decoder := json.NewDecoder(r.Body)
        if err := decoder.Decode(&model); err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }
        defer r.Body.Close()

        if err := updateFunc(db.DB, &model); err != nil {
            http.Error(w, "Failed to update model", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
    }
}

// DeleteHandler handles the deletion of a model by ID
func (db *DBHandler) DeleteHandler(deleteFunc func(*gorm.DB, *domain.Model, string) error) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Query().Get("id")
        var model domain.Model
        if err := deleteFunc(db.DB, &model, id); err != nil {
            http.Error(w, "Failed to delete model", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
    }
}
