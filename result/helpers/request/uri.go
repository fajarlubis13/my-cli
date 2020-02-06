package request

import (
	"fmt"
	"strings"
)

// QueryParameter ...
type QueryParameter struct {
	Search       string
	Limit        int64
	Offset       int64
	SortBy       []string
	ColumnFilter []*ColumnFilter
}

// ColumnFilter ...
type ColumnFilter struct {
	Column   string
	Operator string
	Value    string
	Criteria map[string]string
}

// FormatSort ...
func (q *QueryParameter) FormatSort(firstConcat string) string {
	var (
		sortMap = make(map[string]string, 0)
		order   = []string{"asc", "desc"}

		orderBy string
	)

	if len(q.SortBy) > 1 {
		orderBy += fmt.Sprintf("%s ", firstConcat)
	}

	for _, v := range q.SortBy {
		for _, y := range order {
			if strings.Contains(v, y+"(") {
				column := strings.TrimLeft(strings.TrimRight(v, ")"), y+"(")
				sortMap[column] = strings.ToUpper(y)
			}
		}
	}

	x := 0
	for i, v := range sortMap {
		orderBy += fmt.Sprintf("%s %s", i, v)

		if x%2 == 0 {
			orderBy += ", "
		}
		x++
	}

	return orderBy
}

// AssignColumnFilter ...
func (q *QueryParameter) AssignColumnFilter(m []ColumnFilter) {
	for _, mx := range m {
		for i, v := range mx.Criteria {
			cv := ColumnFilter{}
			cv.Column = mx.Column
			cv.Operator = i
			cv.Value = v

			q.ColumnFilter = append(q.ColumnFilter, &cv)
		}
	}
}

// FormatColumnFilter ...
func (q *QueryParameter) FormatColumnFilter(firstConcat, nextConcat string) string {
	var (
		filter string
	)

	for i, v := range q.ColumnFilter {
		if i != 0 {
			if firstConcat == "WHERE" {
				firstConcat = "AND"
			}
		}

		filter += fmt.Sprintf("%s ", firstConcat)
		filter += fmt.Sprintf("%s %s %s ", v.Column, convertPostgresOperator(v.Operator), v.Value)
	}

	return filter
}

func convertPostgresOperator(input string) string {
	var ops string

	switch input {
	case "eq":
		ops = "="
	case "ne":
		ops = "<>"
	case "gt":
		ops = ">"
	case "gte":
		ops = ">="
	case "lt":
		ops = "<"
	case "lte":
		ops = "<="
	default:
		ops = "="
	}

	return ops
}
