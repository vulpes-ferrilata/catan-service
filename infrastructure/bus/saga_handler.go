package bus

import "context"

type SagaHandler interface {
	GetSaga() interface{}
	Handle(ctx context.Context, data interface{}) error
}
