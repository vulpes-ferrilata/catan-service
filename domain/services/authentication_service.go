package services

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/shared/proto/authentication"
	"github.com/asim/go-micro/v3/client"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type AuthenticationService interface {
	GetClaimByAccessToken(ctx context.Context, accessToken string) (*models.Claim, error)
}

func NewAuthenticationService(client client.Client) AuthenticationService {
	return &authenticationService{
		authenticationService: authentication.NewAuthenticationService("boardgame.authentication.service", client),
	}
}

type authenticationService struct {
	authenticationService authentication.AuthenticationService
}

func (a authenticationService) GetClaimByAccessToken(ctx context.Context, accessToken string) (*models.Claim, error) {
	getClaimByAccessTokenRequest := &authentication.GetClaimByAccessTokenRequest{
		AccessToken: accessToken,
	}

	claimResponse, err := a.authenticationService.GetClaimByAccessToken(ctx, getClaimByAccessTokenRequest)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userID, err := uuid.Parse(claimResponse.GetUserID())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	claim := models.NewClaim(userID)

	return claim, nil
}
