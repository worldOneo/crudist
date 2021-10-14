package gormstorage

import (
	"gorm.io/gorm"
)

// Storage type for gorm
type Storage struct {
	db     *gorm.DB
	config Config
}

// Inceptor for DB lookup
// Return the current query and
// a boolean if the query is done
type Inceptor func(model interface{}, tx *gorm.DB) (*gorm.DB, bool)

// Inceptors to modifie DB lookups otf
// Applied as DefaultLookup(Specific(All())) *gorm.DB
type Inceptors struct {
	// All Always applied inceptors
	All        []Inceptor
	Get        []Inceptor
	GetByID    []Inceptor
	Create     []Inceptor
	Update     []Inceptor
	Delete     []Inceptor
	DeleteByID []Inceptor
}

// Config for additional functionality
type Config struct {
	// Inceptors to modifie DB lookups
	Inceptors Inceptors
}

// Gorm creates a new storage layer
func Gorm(db *gorm.DB, config ...Config) *Storage {
	conf := Config{}
	if len(config) > 0 {
		conf = config[0]
	}
	return &Storage{db, conf}
}

func find(models interface{}, tx *gorm.DB) (*gorm.DB, bool) {
	return tx.Find(models), false
}

func creator(model interface{}, tx *gorm.DB) (*gorm.DB, bool) {
	return tx.Create(model), false
}

func updater(model interface{}, tx *gorm.DB) (*gorm.DB, bool) {
	return tx.Save(model), false
}

func deleter(model interface{}, tx *gorm.DB) (*gorm.DB, bool) {
	return tx.Delete(model), false
}

// Get gets a model slice to populate
func (g *Storage) Get(models interface{}) error {
	return applyInceptorChain(models, g.db, g.config.Inceptors.All, g.config.Inceptors.Get, []Inceptor{find})
}

// GetByID gets a model with the specific id
func (g *Storage) GetByID(model, id interface{}) error {
	finder := func(model interface{}, tx *gorm.DB) (*gorm.DB, bool) {
		return tx.Find(model, id), false
	}
	return applyInceptorChain(model, g.db, g.config.Inceptors.All, g.config.Inceptors.GetByID, []Inceptor{finder})
}

// Create adds a new model
func (g *Storage) Create(model interface{}) error {
	return applyInceptorChain(model, g.db, g.config.Inceptors.All, g.config.Inceptors.Create, []Inceptor{creator})
}

// Update modifies a model
func (g *Storage) Update(model interface{}) error {
	return applyInceptorChain(model, g.db, g.config.Inceptors.All, g.config.Inceptors.Update, []Inceptor{updater})
}

// Delete removes a model
func (g *Storage) Delete(model interface{}) error {
	return applyInceptorChain(model, g.db, g.config.Inceptors.All, g.config.Inceptors.Delete, []Inceptor{deleter})
}

// DeleteByID deletes a model given an id
func (g *Storage) DeleteByID(model, id interface{}) error {
	deleter := func(model interface{}, tx *gorm.DB) (*gorm.DB, bool) {
		return tx.Delete(model, id), false
	}
	return applyInceptorChain(model, g.db, g.config.Inceptors.All, g.config.Inceptors.Delete, []Inceptor{deleter})
}

func applyInceptorChain(model interface{}, db *gorm.DB, inceptorss ...[]Inceptor) error {
	done := false
	for _, inceptors := range inceptorss {
		db, done = applyInceptors(model, db, inceptors)
		if done {
			return db.Error
		}
	}
	return db.Error
}

func applyInceptors(model interface{}, db *gorm.DB, inceptors []Inceptor) (*gorm.DB, bool) {
	done := false
	for _, inceptor := range inceptors {
		db, done = inceptor(model, db)
		if done {
			return db, true
		}
		if db.Error != nil {
			return db, false
		}
	}
	return db, false
}
