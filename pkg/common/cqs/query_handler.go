package cqs

import (
	"context"
	"fmt"
)

// InvalidQueryError is self-described
type InvalidQueryError struct {
	expected string
	had      string
}

// NewInvalidQueryError is a constructor
func NewInvalidQueryError(expected string, had string) InvalidQueryError {
	return InvalidQueryError{expected: expected, had: had}
}

func (e InvalidQueryError) Error() string {
	return fmt.Sprintf("invalid query, expected '%s' but found '%s'", e.expected, e.had)
}

// Query is the interface to identify the DTO for a given query by name.
type Query interface {
	Name() string
}

// QueryName is string to identify a given query when it has not input parameters.
type QueryName string

// Name implements Query interface
func (qn QueryName) Name() string {
	return string(qn)
}

// QueryResult is a generic query result type.
type QueryResult interface{}

// QueryHandler is the interface for handling queries.
type QueryHandler interface {
	Handle(ctx context.Context, query Query) (QueryResult, error)
}

type queryHandlerFunc func(ctx context.Context, query Query) (QueryResult, error)

func (f queryHandlerFunc) Handle(ctx context.Context, query Query) (QueryResult, error) {
	return f(ctx, query)
}
