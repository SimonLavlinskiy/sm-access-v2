package queryService

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Queries struct {
	Page    int
	Per     int
	Deleted string
	Name    string
}

type Meta struct {
	Page       int `json:"page"`
	Per        int `json:"per"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

func GetQueries(c *gin.Context) (Queries, Meta) {
	var page int
	var per int
	var deleted string
	var name string

	if q := c.Query("page"); len(q) == 0 {
		page = 1
	} else {
		page, _ = strconv.Atoi(q)
	}

	if q := c.Query("per"); len(q) == 0 {
		per = 30
	} else {
		per, _ = strconv.Atoi(q)
	}

	if q := c.Query("deleted"); len(q) != 0 {
		deleted = q
	}

	if q := c.Query("name"); len(q) != 0 {
		name = q
	}

	queries := Queries{
		Page:    page,
		Per:     per,
		Deleted: deleted,
		Name:    name,
	}
	meta := Meta{
		Page: page,
		Per:  per,
	}

	return queries, meta
}
