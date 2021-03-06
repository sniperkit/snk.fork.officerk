package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
)

// JobType has two types
type JobType int

// JobType ...
const (
	JobTypeCron JobType = iota
	JobTypeManual
)

// Job contains at least one task
type Job struct {
	Comm
	Name     string     `gorm:"not null" json:"name"`
	Typ      JobType    `gorm:"not null" json:"typ"`
	Schedule NullString `json:"schedule,omitempty"`
	Slug     NullString `gorm:"unique" json:"slug,omitempty"`
	IsOnline bool       `gorm:"not null" json:"is_online"`
	NodeID   uint       `gorm:"not null" json:"node_id"`

	Node  Node   `json:"node"`
	Tasks []Task `gorm:"ForeignKey:JobID" json:"tasks,omitempty"`
}
