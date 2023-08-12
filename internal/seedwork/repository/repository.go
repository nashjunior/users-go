package repository

import "school-api/internal/seedwork/entities"

type OrderSort string

const (
	Asc  OrderSort = "asc"
	Desc OrderSort = "desc"
)

type Sort struct {
	Field    string
	SortType OrderSort
}

type SearchProps[Filter any] struct {
	Page       *int
	PerPage    *int
	SortFields *[]string
	OrderSort  *[]OrderSort
	Filter     *Filter
}

type Counter[T any] struct {
	total int
	items []T
}

type Repository[T any] interface {
	Create(instance T) error
	CreateBatch(instance []T) error
	FindById(id string) (*T, error)
	find() []T
	findAndCount() (Counter[T], error)
	update(data T) error
	delete(id string) error
}

type SearchRepository[
	T entities.AbstractEntity,
	Filter any,
	SearchableParams SearchParams[Filter],
	Res SearchResult[T, Filter]] interface {
	search(props SearchableParams) Res
}
