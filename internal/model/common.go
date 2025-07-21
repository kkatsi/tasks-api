package model

import "fmt"

type PaginationParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (p *PaginationParams) Validate() error {

	if p.Limit <= 0 {
		return fmt.Errorf("limit must be greater than 0")
	}
	if p.Limit > 100 {
		return fmt.Errorf("limit cannot exceed 100")
	}
	if p.Offset < 0 {
		return fmt.Errorf("offset cannot be negative")
	}

	return nil
}

// Helper to get validated values
func (p *PaginationParams) GetValues() (limit, offset int) {
	return p.Limit, p.Offset
}

func (p *PaginationParams) GetPage() int {
	if p.Limit == 0 {
		return 1
	}
	return (p.Offset / p.Limit) + 1
}
