package bus

import (
	"context"
	"reflect"
	"sync"

	"github.com/pkg/errors"
)

type QueryBus interface {
	Register(queryHandlers ...QueryHandler) error
	Use(handlerWrapper ...HandlerWrapper)
	Execute(ctx context.Context, query interface{}) (interface{}, error)
}

func NewQueryBus() QueryBus {
	queryBus := &queryBus{
		handlers:        make(map[string]QueryHandler),
		handlerWrappers: make([]HandlerWrapper, 0),
	}

	return queryBus
}

type queryBus struct {
	handlers        map[string]QueryHandler
	handlerWrappers []HandlerWrapper
	mu              sync.RWMutex
}

func (q *queryBus) getQueryHandler(query interface{}) (QueryHandler, error) {
	queryName := reflect.TypeOf(query).String()

	q.mu.RLock()
	defer q.mu.RUnlock()
	handler, ok := q.handlers[queryName]
	if !ok {
		return nil, errors.Errorf("handler not found for query (%s)", queryName)
	}

	return handler, nil
}

func (q *queryBus) setQueryHandler(queryHandler QueryHandler) error {
	query := queryHandler.GetQuery()

	queryName := reflect.TypeOf(query).String()

	if handler, err := q.getQueryHandler(query); err == nil {
		return errors.Errorf("query (%s) is already assigned to handler (%s)", queryName, reflect.TypeOf(handler))
	}

	q.mu.Lock()
	defer q.mu.Unlock()
	q.handlers[queryName] = queryHandler

	return nil
}

func (q *queryBus) Register(queryHandlers ...QueryHandler) error {
	for _, queryHandler := range queryHandlers {
		if err := q.setQueryHandler(queryHandler); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (q *queryBus) Use(handlerWrapper ...HandlerWrapper) {
	q.handlerWrappers = append(q.handlerWrappers, handlerWrapper...)
}

func (q *queryBus) Execute(ctx context.Context, query interface{}) (interface{}, error) {
	queryHandler, err := q.getQueryHandler(query)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	fn := func(ctx context.Context, input interface{}) (interface{}, error) {
		return queryHandler.Handle(ctx, input)
	}

	for i := len(q.handlerWrappers) - 1; i >= 0; i-- {
		fn = q.handlerWrappers[i](fn)
	}

	return fn(ctx, query)
}
