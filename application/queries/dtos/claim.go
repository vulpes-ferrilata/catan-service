package dtos

import "github.com/VulpesFerrilata/catan-service/domain/models"

func NewClaimDTO(claim *models.Claim) *ClaimDTO {
	return &ClaimDTO{
		UserID: claim.GetUserID().String(),
	}
}

type ClaimDTO struct {
	UserID string
}
