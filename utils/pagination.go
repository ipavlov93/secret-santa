package utils

import (
	"fmt"
)

const (
	defaultPage = 1

	defaultLimit = 10
	maxLimit     = 1000
)

type Pagination struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}

func NewPagination(page uint64) Pagination {
	if page == 0 {
		page = defaultPage
	}
	return Pagination{
		Page:  page,
		Limit: defaultLimit,
	}
}

func (p *Pagination) Validate(itemsCounter uint64) error {
	if p.Page == 0 {
		p.Page = defaultPage
	}
	if p.Limit == 0 {
		p.Limit = defaultLimit
	}
	if p.Limit > maxLimit {
		p.Limit = maxLimit
	}
	if p.Offset()+p.Limit > itemsCounter {
		return fmt.Errorf("no more data exist")
	}
	return nil
}

func (p Pagination) Offset() (offset uint64) {
	if p.Page == defaultPage {
		return 0
	}
	return (p.Page - 1) * p.Limit
}
