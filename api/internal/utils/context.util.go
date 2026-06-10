package utils

import (
	"context"

	"go-cover-parroto/internal/core/response"
)

func GetFromContext[T any](ctx context.Context, key any) (T, *response.AppError) {
	val, ok := ctx.Value(key).(T)
	if !ok {
		var zero T
		return zero, response.Unauthorized("missing context value")
	}
	return val, nil
}
