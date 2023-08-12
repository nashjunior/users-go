package test

import (
	"reflect"
	"school-api/internal/seedwork/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldHandlePageProp(t *testing.T) {
	arrange := []struct {
		page     interface{}
		expected int
	}{
		{page: nil, expected: 1},
		{page: -1, expected: 1},
		{page: "", expected: 1},
		{page: true, expected: 1},
		{page: 5.1, expected: 5},
	}

	for _, v := range arrange {

		pageValue := 0
		if v.page != nil {
			switch v.page.(type) {
			case int:
				pageValue = v.page.(int)
			case float64:
				pageValue = int(v.page.(float64))

			default:
				pageValue = 1 // Default value if type assertion fails
			}
		}

		params := repository.NewSearchParams[any](repository.SearchProps[any]{
			Page: &pageValue,
		})

		assert.Equal(t, v.expected, params.GetPage())
	}

}

func TestShouldHandlePerPageProp(t *testing.T) {
	arrange := []struct {
		perPage  interface{}
		expected int
	}{
		{perPage: nil, expected: 10},
		{perPage: "", expected: 10},
		{perPage: -1, expected: 10},
		{perPage: true, expected: 10},
		{perPage: 5.1, expected: 5},
	}

	for _, v := range arrange {

		perPageValue := 0
		if v.perPage != nil {
			switch v.perPage.(type) {
			case int:
				perPageValue = v.perPage.(int)
			case float64:
				perPageValue = int(v.perPage.(float64))

			default:
				perPageValue = 0 // Default value if type assertion fails
			}
		}

		params := repository.NewSearchParams[any](repository.SearchProps[any]{
			PerPage: &perPageValue,
		})

		assert.Equal(t, v.expected, params.GetPerPage())
	}

}

func TestShouldHandleSortProp(t *testing.T) {
	arrange := []struct {
		sort     interface{}
		expected *[]repository.Sort
	}{
		{sort: nil, expected: &[]repository.Sort{}},
		{sort: []string{}, expected: &[]repository.Sort{}},
		{sort: []string{"asd"}, expected: &[]repository.Sort{
			{Field: "asd", SortType: repository.Asc},
		}},
	}

	for _, v := range arrange {
		var sortValue *[]string

		if v.sort != nil && reflect.TypeOf(v.sort).Kind() == reflect.Slice {
			arr, ok := v.sort.([]string)
			if ok {
				sortValue = &arr
			}
		}

		params := repository.NewSearchParams[any](repository.SearchProps[any]{
			SortFields: sortValue,
			OrderSort:  &[]repository.OrderSort{repository.Asc},
		})

		assert.Equal(t, v.expected, params.GetSort())
	}
}

func TestShouldHandleFilterProp(t *testing.T) {
	arrange := []struct {
		filter   *interface{}
		expected *any
	}{
		{filter: toInterface(nil), expected: toInterface(nil)},
		{filter: toInterface(-1), expected: toInterface(-1)},
		{filter: toInterface("-1"), expected: toInterface("-1")},
		{filter: toInterface([]interface{}{}), expected: toInterface([]interface{}{})},
	}

	for _, v := range arrange {

		params := repository.NewSearchParams[any](repository.SearchProps[any]{
			Filter: v.filter,
		})

		response := params.GetFilter()
		assert.Equal(t, *v.expected, *response)
	}
}

func toInterface(value interface{}) *interface{} {
	return &value
}
