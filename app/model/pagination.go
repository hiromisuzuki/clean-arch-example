package model

import "strconv"

//Pagination Pagination
type Pagination struct {
	Page     int `json:"page" validate:"number, gte=1,lte=999" example:"1"`
	PerPage  int `json:"per_page" validate:"number, gte=1,lte=100" example:"1"`
	MaxPages int `json:"max_page" validate:"number" example:"1"`
}

//Limit Limit
func (p *Pagination) Limit() string {
	return strconv.Itoa(p.PerPage)
}

//Offset Offset
func (p *Pagination) Offset() string {
	return strconv.Itoa((p.Page - 1) * p.PerPage)
}
