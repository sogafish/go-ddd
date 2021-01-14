package models

import (
  "time"

  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Item struct {
  ID           int `gorm:"primaryKey" json:"id"`
  Title        string     `gorm:"size:255" json:"title,omitempty"`
  CreatedAt    *time.Time `json:"created_at"`
  UpdatedAt    *time.Time `json:"updated_at"`
}
