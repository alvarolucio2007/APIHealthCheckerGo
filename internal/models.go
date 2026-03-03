// Package internal Contém a estrutura de dados e a persistencia do monitor.
package internal

import (
	"time"

	"gorm.io/gorm"
)

type ChecagemAPI struct {
	gorm.Model
	LinkAPI    string `gorm:"not null"`
	IsOnline   bool
	StatusCode int
	Latency    time.Duration
}
