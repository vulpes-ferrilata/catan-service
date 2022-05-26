package saga

import (
	"context"

	"github.com/pkg/errors"
)

type Coordinator interface {
	WithStep(step *Step) Coordinator
	Execute(ctx context.Context) error
}

func NewCoordinator() Coordinator {
	return &coordinator{
		steps: make([]*Step, 0),
	}
}

type coordinator struct {
	steps []*Step
}

func (s *coordinator) WithStep(step *Step) Coordinator {
	s.steps = append(s.steps, step)

	return s
}

func (s coordinator) handle(ctx context.Context) error {
	for i := 0; i < len(s.steps); i++ {
		if err := s.steps[i].Handle(ctx); err != nil {
			if err := s.compensate(ctx, i-1); err != nil {
				return errors.WithStack(err)
			}

			return errors.WithStack(err)
		}
	}

	return nil
}

func (s coordinator) compensate(ctx context.Context, idx int) error {
	for i := idx; i >= 0; i-- {
		if err := s.steps[i].Compensate(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (s coordinator) Execute(ctx context.Context) error {
	if err := s.handle(ctx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
