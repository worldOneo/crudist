package gormstorage

import "gorm.io/gorm"

// Storage type for gorm
type Storage struct {
	db *gorm.DB
}

func Gorm(db *gorm.DB) *Storage {
	return &Storage{db}
}

// Get gets a model slice to populate
func (g *Storage) Get(models interface{}) error {
	return g.db.Find(models).Error
}

// GetByID gets a model with the specific id
func (g *Storage) GetByID(model interface{}, id interface{}) error {
	return g.db.Find(model, id).Error
}

// Create adds a new model
func (g *Storage) Create(model interface{}) error {
	return g.db.Create(model).Error
}

// Update modifies a model
func (g *Storage) Update(model interface{}) error {
	return g.db.Save(model).Error
}

// Delete removes a model
func (g *Storage) Delete(model interface{}) error {
	return g.db.Delete(model).Error
}
