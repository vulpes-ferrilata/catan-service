package errors

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAuthenticationError(err error) AppError {
	if detailErr, ok := errors.Cause(err).(DetailError); ok {
		err = NewBusinessRuleError(detailErr)
	}

	return &authenticationError{
		err: err,
	}
}

type authenticationError struct {
	err error
}

func (a authenticationError) Error() string {
	builder := new(strings.Builder)

	builder.WriteString("authentication failed")
	builder.WriteString("\n")
	builder.WriteString(a.err.Error())

	return builder.String()
}

func (a authenticationError) Problem(translator ut.Translator) (iris.Problem, error) {
	problem := iris.NewProblem()
	var err error

	if webErr, ok := errors.Cause(a.err).(WebError); ok {
		problem, err = webErr.Problem(translator)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	} else {
		problem.Detail(a.err.Error())
	}

	problem.Status(iris.StatusUnauthorized)

	title, err := translator.T("authentication-error-title")
	if err != nil {
		return nil, errors.Wrap(err, "authentication-error-title")
	}
	problem.Title(title)

	return problem, nil
}

func (a authenticationError) Status(translator ut.Translator, serviceName string) (*status.Status, error) {
	stt := status.New(codes.FailedPrecondition, serviceName)

	if grpcErr, ok := errors.Cause(a.err).(GrpcError); ok {
		stt, err := grpcErr.Status(translator, serviceName)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		sttProto := stt.Proto()
		sttProto.Code = int32(codes.FailedPrecondition)
		stt = status.FromProto(sttProto)
	} else {
		sttProto := stt.Proto()
		sttProto.Message = a.err.Error()
		stt = status.FromProto(sttProto)
	}

	return stt, nil
}
