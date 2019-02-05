package controller

import (
	"strconv"

	"github.com/hiromisuzuki/clean-arch-example/app/model"
)

//NewPagination NewPagination
func NewPagination(page string, perPage string) model.Pagination {
	return model.Pagination{
		Page:    parseInt(page, 1),
		PerPage: parseInt(perPage, 100),
	}
}

func parseInt(s string, sub int) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		res = sub
	}
	return res
}
