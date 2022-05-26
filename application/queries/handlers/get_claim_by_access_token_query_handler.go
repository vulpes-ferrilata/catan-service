package handlers

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/application/queries"
	"github.com/VulpesFerrilata/catan-service/application/queries/dtos"
	"github.com/VulpesFerrilata/catan-service/domain/services"
	"github.com/VulpesFerrilata/catan-service/infrastructure/dig/results"
	"github.com/pkg/errors"
)

func NewGetClaimByAccessTokenQueryHandler(authenticationService services.AuthenticationService) results.QueryHandlerResult {
	queryHandler := &getClaimByAccessTokenQueryHandler{
		authenticationService: authenticationService,
	}

	return results.QueryHandlerResult{
		QueryHandler: queryHandler,
	}
}

type getClaimByAccessTokenQueryHandler struct {
	authenticationService services.AuthenticationService
}

func (g getClaimByAccessTokenQueryHandler) GetQuery() interface{} {
	return new(queries.GetClaimByAccessTokenQuery)
}

func (g getClaimByAccessTokenQueryHandler) Handle(ctx context.Context, query interface{}) (interface{}, error) {
	getClaimByAccessTokenQuery := query.(*queries.GetClaimByAccessTokenQuery)

	claim, err := g.authenticationService.GetClaimByAccessToken(ctx, getClaimByAccessTokenQuery.AccessToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	claimDTO := dtos.NewClaimDTO(claim)

	return claimDTO, nil
}
