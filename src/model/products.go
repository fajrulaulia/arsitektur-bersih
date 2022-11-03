package model

import "time"

type Product struct {
	Code      string
	Name      string
	Price     float64
	CreatetAt time.Time
	UpdatedAt time.Time
}
