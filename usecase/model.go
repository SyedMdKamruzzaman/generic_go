package usecase

import (
    "github.com/jinzhu/gorm"
    "mes/domain"
)

// DBHandler contains the database connection
type DBHandler struct {
    DB *gorm.DB
}

// NewDBHandler creates a new DBHandler instance
func NewDBHandler(db *gorm.DB) *DBHandler {
    return &DBHandler{DB: db}
}

// Create generates a new model
func (handler *DBHandler) Create(model *domain.Model) error {
    if err := handler.DB.Create(model).Error; err != nil {
        return err
    }
    return nil
}

// Read retrieves a model by ID
func (handler *DBHandler) Read(model *domain.Model, id string) error {
    if err := handler.DB.First(model, id).Error; err != nil {
        return err
    }
    return nil
}

// ReadAll retrieves all records of a model from the database
func (handler *DBHandler) ReadAll(models *[]domain.Model) error {
    if err := handler.DB.Find(models).Error; err != nil {
        return err
    }
    return nil
}

// Update updates an existing model
func (handler *DBHandler) Update(model *domain.Model) error {
    if err := handler.DB.Save(model).Error; err != nil {
        return err
    }
    return nil
}

// Delete deletes a model by ID
func (handler *DBHandler) Delete(model *domain.Model, id string) error {
    if err := handler.DB.Delete(model, id).Error; err != nil {
        return err
    }
    return nil
}
