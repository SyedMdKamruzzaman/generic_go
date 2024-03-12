// File: usecase/model.go

package usecase

import (
    "github.com/jinzhu/gorm"
    "mes/domain"
)

// Create generates a new model
func Create(db *gorm.DB, model *domain.Model) error {
    if err := db.Create(model).Error; err != nil {
        return err
    }
    return nil
}

// Read retrieves a model by ID
func Read(db *gorm.DB, model *domain.Model, id string) error {
    if err := db.First(model, id).Error; err != nil {
        return err
    }
    return nil
}

// ReadAll retrieves all records of a model from the database
func ReadAll(db *gorm.DB, models *[]domain.Model) error {
    if err := db.Find(models).Error; err != nil {
        return err
    }
    return nil
}

// Update updates an existing model
func Update(db *gorm.DB, model *domain.Model) error {
    if err := db.Save(model).Error; err != nil {
        return err
    }
    return nil
}

// Delete deletes a model by ID
func Delete(db *gorm.DB, model *domain.Model, id string) error {
    if err := db.Delete(model, id).Error; err != nil {
        return err
    }
    return nil
}
