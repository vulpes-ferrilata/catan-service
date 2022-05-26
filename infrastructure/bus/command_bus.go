package bus

import (
	"context"
	"reflect"
	"sync"

	"github.com/pkg/errors"
)

type CommandBus interface {
	Register(commandHandlers ...CommandHandler) error
	Use(handlerWrappers ...HandlerWrapper)
	Execute(ctx context.Context, command interface{}) error
}

func NewCommandBus() CommandBus {
	commandBus := &commandBus{
		handlers:        make(map[string]CommandHandler),
		handlerWrappers: make([]HandlerWrapper, 0),
	}

	return commandBus
}

type commandBus struct {
	handlers        map[string]CommandHandler
	handlerWrappers []HandlerWrapper
	mu              sync.RWMutex
}

func (c *commandBus) getCommandHandler(command interface{}) (CommandHandler, error) {
	commandName := reflect.TypeOf(command).String()

	c.mu.RLock()
	defer c.mu.RUnlock()
	handler, ok := c.handlers[commandName]
	if !ok {
		return nil, errors.Errorf("handler not found for command (%s)", commandName)
	}

	return handler, nil
}

func (c *commandBus) setCommandHandler(commandHandler CommandHandler) error {
	command := commandHandler.GetCommand()

	commandName := reflect.TypeOf(command).String()

	if handler, err := c.getCommandHandler(command); err == nil {
		return errors.Errorf("command (%s) is already assigned to handler (%s)", commandName, reflect.TypeOf(handler))
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers[commandName] = commandHandler

	return nil
}

func (c *commandBus) Register(commandHandlers ...CommandHandler) error {
	for _, commandHandler := range commandHandlers {
		if err := c.setCommandHandler(commandHandler); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (c *commandBus) Use(handlerWrappers ...HandlerWrapper) {
	c.handlerWrappers = append(c.handlerWrappers, handlerWrappers...)
}

func (c *commandBus) Execute(ctx context.Context, command interface{}) error {
	commandHandler, err := c.getCommandHandler(command)
	if err != nil {
		return errors.WithStack(err)
	}

	fn := func(ctx context.Context, input interface{}) (interface{}, error) {
		if err := commandHandler.Handle(ctx, input); err != nil {
			return nil, errors.WithStack(err)
		}

		return nil, nil
	}

	for i := len(c.handlerWrappers) - 1; i >= 0; i-- {
		fn = c.handlerWrappers[i](fn)
	}

	if _, err := fn(ctx, command); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
