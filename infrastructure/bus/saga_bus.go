package bus

import (
	"context"
	"reflect"
	"sync"

	"github.com/pkg/errors"
)

type SagaBus interface {
	Register(sagaHandlers ...SagaHandler) error
	Use(handlerWrappers ...HandlerWrapper)
	Execute(ctx context.Context, saga interface{}) error
}

func NewSagaBus() SagaBus {
	return &sagaBus{
		handlers:        make(map[string]SagaHandler),
		handlerWrappers: make([]HandlerWrapper, 0),
	}
}

type sagaBus struct {
	handlers        map[string]SagaHandler
	handlerWrappers []HandlerWrapper
	mu              sync.RWMutex
}

func (s *sagaBus) getSagaHandler(data interface{}) (SagaHandler, error) {
	dataName := reflect.TypeOf(data).String()

	s.mu.RLock()
	defer s.mu.RUnlock()
	handler, ok := s.handlers[dataName]
	if !ok {
		return nil, errors.Errorf("handler not found for data (%s)", dataName)
	}

	return handler, nil
}

func (s *sagaBus) setSagaHandler(sagaHandler SagaHandler) error {
	saga := sagaHandler.GetSaga()

	sagaName := reflect.TypeOf(saga).String()

	if handler, err := s.getSagaHandler(saga); err == nil {
		return errors.Errorf("saga (%s) is already assigned to handler (%s)", sagaName, reflect.TypeOf(handler))
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[sagaName] = sagaHandler

	return nil
}

func (s *sagaBus) Register(sagaHandlers ...SagaHandler) error {
	for _, sagaHandler := range sagaHandlers {
		if err := s.setSagaHandler(sagaHandler); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (s *sagaBus) Use(handlerWrappers ...HandlerWrapper) {
	s.handlerWrappers = append(s.handlerWrappers, handlerWrappers...)
}

func (s *sagaBus) Execute(ctx context.Context, saga interface{}) error {
	sagaHandler, err := s.getSagaHandler(saga)
	if err != nil {
		return errors.WithStack(err)
	}

	fn := func(ctx context.Context, input interface{}) (interface{}, error) {
		if err := sagaHandler.Handle(ctx, input); err != nil {
			return nil, errors.WithStack(err)
		}

		return nil, nil
	}

	for i := len(s.handlerWrappers) - 1; i >= 0; i-- {
		fn = s.handlerWrappers[i](fn)
	}

	if _, err := fn(ctx, saga); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
