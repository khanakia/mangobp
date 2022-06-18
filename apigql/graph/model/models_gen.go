// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type FilterInput struct {
	Condition *FilterCondition `json:"condition"`
	Field     string           `json:"field"`
	Type      *string          `json:"type"`
	Operator  *FilterOperator  `json:"operator"`
	Value     interface{}      `json:"value"`
}

type MetaInput struct {
	Limit   int               `json:"limit"`
	OrderBy []*SortOrderInput `json:"orderBy"`
	Offset  int               `json:"offset"`
}

type MetaOutput struct {
	Total *int `json:"total"`
}

type SortOrderInput struct {
	Key   string         `json:"key"`
	Value *SortDirection `json:"value"`
}

type SortOrderOutput struct {
	Key   *string        `json:"key"`
	Value *SortDirection `json:"value"`
}

// The basic FilterCondition
type FilterCondition string

const (
	FilterConditionAnd FilterCondition = "AND"
	FilterConditionOr  FilterCondition = "OR"
)

var AllFilterCondition = []FilterCondition{
	FilterConditionAnd,
	FilterConditionOr,
}

func (e FilterCondition) IsValid() bool {
	switch e {
	case FilterConditionAnd, FilterConditionOr:
		return true
	}
	return false
}

func (e FilterCondition) String() string {
	return string(e)
}

func (e *FilterCondition) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterCondition(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterCondition", str)
	}
	return nil
}

func (e FilterCondition) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FilterOperator string

const (
	FilterOperatorLike    FilterOperator = "LIKE"
	FilterOperatorEqualto FilterOperator = "EQUALTO"
)

var AllFilterOperator = []FilterOperator{
	FilterOperatorLike,
	FilterOperatorEqualto,
}

func (e FilterOperator) IsValid() bool {
	switch e {
	case FilterOperatorLike, FilterOperatorEqualto:
		return true
	}
	return false
}

func (e FilterOperator) String() string {
	return string(e)
}

func (e *FilterOperator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterOperator", str)
	}
	return nil
}

func (e FilterOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The basic directions
type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)

var AllSortDirection = []SortDirection{
	SortDirectionAsc,
	SortDirectionDesc,
}

func (e SortDirection) IsValid() bool {
	switch e {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	}
	return false
}

func (e SortDirection) String() string {
	return string(e)
}

func (e *SortDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirection", str)
	}
	return nil
}

func (e SortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}