package repository

import (
	"fmt"
	"math"
	"strings"
)

type SearchParams[Filter any] struct {
	page    int
	perPage int
	sort    *[]Sort
	filter  *Filter
}

func NewSearchParams[T any](props SearchProps[T]) *SearchParams[T] {
	params := &SearchParams[T]{}

	params.SetPage(props.Page)
	params.SetPerPage(props.PerPage)
	params.SetSort(props.SortFields, props.OrderSort)
	params.SetFilter(props.Filter)

	return params
}

func (params *SearchParams[Filter]) SetPage(page *int) {
	if page != nil && !math.IsNaN(float64(*page)) && *page > 0 {
		params.page = *page
	} else {
		params.page = 1
	}
}

func (params *SearchParams[Filter]) GetPage() int {
	return params.page
}

func (params *SearchParams[Filter]) SetPerPage(perPage *int) {
	if perPage != nil && !math.IsNaN(float64(*perPage)) && *perPage > 0 {
		params.perPage = *perPage
	} else {
		params.perPage = 10
	}
}

func (params *SearchParams[Filter]) GetPerPage() int {
	return params.perPage
}

func (params *SearchParams[Filter]) SetSort(sort *[]string, orderSort *[]OrderSort) {
	if sort == nil || orderSort == nil || len(*sort) == 0 || len(*orderSort) == 0 {
		params.sort = &[]Sort{}
		return
	}

	if len(*sort) != len(*orderSort) {
		params.sort = &[]Sort{}
		return
	}

	var filteredSort []Sort
	for index, field := range *sort {
		order := (*orderSort)[index]
		if order == Asc || order == Desc {
			filteredSort = append(filteredSort, Sort{
				Field:    field,
				SortType: order,
			})
		}
	}

	params.sort = &filteredSort
}

func (params *SearchParams[Filter]) GetSort() *[]Sort {
	return params.sort
}

func (params *SearchParams[Filter]) SetFilter(filter *Filter) {

	var filterString *string

	if filter != nil {
		filteedString := fmt.Sprint(filter)
		filterString = &filteedString
	}

	if filterString != nil && len(strings.TrimSpace(*filterString)) != 0 {
		params.filter = filter
	}

}
func (params *SearchParams[Filter]) GetFilter() *Filter {
	return params.filter
}
