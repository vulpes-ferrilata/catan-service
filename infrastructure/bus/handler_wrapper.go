package bus

import "context"

type HandlerFunc func(context.Context, interface{}) (interface{}, error)

type HandlerWrapper func(next HandlerFunc) HandlerFunc
