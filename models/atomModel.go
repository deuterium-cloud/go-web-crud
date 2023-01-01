package models

import (
	"gorm.io/gorm"
)

type Atom struct {
	gorm.Model
	AtomNumber uint
	Mass       float64
	Name       string
	Symbol     string
}

type AtomRequest struct {
	AtomNumber uint    `json:"number"`
	Mass       float64 `json:"mass"`
	Name       string  `json:"name"`
	Symbol     string  `json:"symbol"`
}

type ResponseDto struct {
	Atoms *[]Atom `json:"atoms"`
	Total int64   `json:"total"`
	Page  int64   `json:"page"`
	Size  int64   `json:"size"`
}

func Map(dto AtomRequest) Atom {
	return Atom{
		AtomNumber: dto.AtomNumber,
		Mass:       dto.Mass,
		Name:       dto.Name,
		Symbol:     dto.Symbol,
	}
}
