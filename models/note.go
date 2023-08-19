package models

import (
	"gorm.io/gorm"
	"time"
)

type Note struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"index"`
	DeletedAt gorm.DeletedAt
}

func NotesAll() *[]Note {
	var notes []Note
	DB.Where("deleted_at is NULL").Order("updated_at desc").Find(&notes)
	return &notes
}

func NoteCreate(name string, content string) *Note {
	entry := Note{Name: name, Content: content}
	DB.Create(&entry)
	return &entry
}

func NoteFind(id uint64) *Note {
	var note Note
	DB.Where("id = ?", id).First(&note)
	return &note
}
