package cqs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// AppError is a query/command hnd error with context
type AppError struct {
	Name   string      `json:"name"`
	Input  interface{} `json:"input"`
	ErrMsg string      `json:"err_msg"`
}

// QueryHandlerMiddleware is a type for decorating QueryHandlers
type QueryHandlerMiddleware func(h QueryHandler) QueryHandler

// NewQueryHndErrorMiddleware is a middleware constructor to log a contextualized query handler error
func NewQueryHndErrorMiddleware(logger *log.Logger) QueryHandlerMiddleware {
	return func(h QueryHandler) QueryHandler {
		return queryHandlerFunc(func(ctx context.Context, q Query) (QueryResult, error) {
			result, err := h.Handle(ctx, q)
			if err != nil {
				logAppErr(logger, AppError{
					Name:   q.Name(),
					Input:  q,
					ErrMsg: err.Error(),
				})
				return nil, err
			}

			return result, nil
		})
	}
}

func logAppErr(logger *log.Logger, appErr AppError) {
	b, err := json.Marshal(&appErr)
	if err != nil {
		msg := fmt.Sprintf("something when wrong when trying to marshal app error from %s: %s", err.Error(), appErr.Name)
		logger.Println([]byte(msg))
		return
	}

	logger.Println(string(b))
}

type CommandHandlerMiddleware func(h CommandHandler) CommandHandler

func NewCommandHndErrorMiddleware(logger *log.Logger) CommandHandlerMiddleware {
	return func(h CommandHandler) CommandHandler {
		return CommandHandlerFunc(func(ctx context.Context, q Command) ([]Event, error) {
			result, err := h.Handle(ctx, q)
			if err != nil {
				logAppErr(logger, AppError{
					Name:   q.Name(),
					Input:  q,
					ErrMsg: err.Error(),
				})
				return nil, err
			}

			return result, nil
		})
	}
}
