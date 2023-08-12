package repository

import (
	"math"
	"school-api/internal/seedwork/entities"
)

type SearchResultProps[E entities.AbstractEntity, Filter any] struct {
	Items       []E
	Total       int
	CurrentPage int
	PerPage     int
	Sort        *[]string
	OrderSort   *[]OrderSort
	Filter      *Filter
}

type SearchResult[Entity entities.AbstractEntity, Filter any] struct {
	Items       []Entity
	Total       int
	CurrentPage int
	PerPage     int
	Lastpage    int
	Sort        *[]string
	OrderSort   *[]OrderSort
	Filter      *Filter
}

func NewSearchResult[E entities.AbstractEntity, Filter any](props SearchResultProps[E, Filter]) SearchResult[E, Filter] {

	return SearchResult[E, Filter]{
		Items:       props.Items,
		Total:       props.Total,
		CurrentPage: props.CurrentPage,
		PerPage:     props.PerPage,
		Lastpage:    int(math.Ceil(float64(props.Total) / float64(props.PerPage))),
		Sort:        props.Sort,
		OrderSort:   props.OrderSort,
	}

}
