package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	ProjectId uint
	Project   Project
	Message   string
}

// SaveNewEntry: save a new entry to the db
func (p *Project) SaveNewEntry(message string, db *gorm.DB) {
	db.Create(&Entry{Message: message, ProjectId: p.ID})
}

// DeleteEntry: delete an entry by id
func DeleteEntry(pKey int, db *gorm.DB) {
	fmt.Println(pKey)
	db.Delete(&Entry{}, pKey)
}

// GetEntriesByDate: return all entries in a date range
func GetEntriesByDate(start time.Time, end time.Time, db *gorm.DB) []Entry {
	var entries []Entry
	db.Where("created_at >= ? and created_at <= ?", start, end).Find(&entries)
	return entries
}
