package middlewares

import (
	"strings"

	service "github.com/VulpesFerrilata/catan-service"
	"github.com/VulpesFerrilata/catan-service/application/queries"
	"github.com/VulpesFerrilata/catan-service/application/queries/dtos"
	"github.com/VulpesFerrilata/catan-service/infrastructure/bus"
	infrastructure_context "github.com/VulpesFerrilata/catan-service/infrastructure/context"
	app_errorss "github.com/VulpesFerrilata/catan-service/infrastructure/errors"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
)

type TokenExtractor func(ctx iris.Context) (string, error)

func FromAuthHeader(ctx iris.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.WithStack(service.ErrHeaderFormatIsInvalid)
	}

	return authHeaderParts[1], nil
}

func FromParameter(ctx iris.Context) (string, error) {
	return ctx.URLParam("token"), nil
}

func FromFirst(tokenExtractors ...TokenExtractor) TokenExtractor {
	return func(ctx iris.Context) (string, error) {
		for _, tokenExtractor := range tokenExtractors {
			accessToken, err := tokenExtractor(ctx)
			if err != nil {
				return "", errors.WithStack(err)
			}
			if accessToken != "" {
				return accessToken, nil
			}
		}

		return "", nil
	}
}

func NewAuthenticationMiddleware(queryBus bus.QueryBus,
	errorHandlerMiddleware *ErrorHandlerMiddleware) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		queryBus:               queryBus,
		errorHandlerMiddleware: errorHandlerMiddleware,
	}
}

type AuthenticationMiddleware struct {
	queryBus               bus.QueryBus
	errorHandlerMiddleware *ErrorHandlerMiddleware
}

func (a AuthenticationMiddleware) handleError(ctx iris.Context, err error) {
	err = app_errorss.NewAuthenticationError(err)
	a.errorHandlerMiddleware.Handle(ctx, err)
}

func (a AuthenticationMiddleware) Serve(ctx iris.Context) {
	accessToken, err := FromFirst(FromAuthHeader, FromParameter)(ctx)
	if err != nil {
		a.handleError(ctx, err)
		return
	}

	request := ctx.Request()
	requestCtx := request.Context()

	getClaimByAccessTokenQuery := &queries.GetClaimByAccessTokenQuery{
		AccessToken: accessToken,
	}
	result, err := a.queryBus.Execute(ctx.Request().Context(), getClaimByAccessTokenQuery)
	if err != nil {
		a.handleError(ctx, err)
		return
	}
	claimDTO := result.(*dtos.ClaimDTO)

	requestCtx = infrastructure_context.WithUserID(requestCtx, claimDTO.UserID)
	request = request.WithContext(requestCtx)
	ctx.ResetRequest(request)

	ctx.Next()
}
