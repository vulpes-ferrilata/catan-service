package context

import (
	"context"

	"github.com/pkg/errors"
)

type userIDContextKey struct{}

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDContextKey{}, userID)
}

func GetUserID(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(userIDContextKey{}).(string)
	if !ok {
		return "", errors.Wrap(ErrContextValueNotFound, "userID")
	}

	return userID, nil
}
