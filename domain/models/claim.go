package models

import "github.com/google/uuid"

func NewClaim(userID uuid.UUID) *Claim {
	return &Claim{
		userID: userID,
	}
}

type Claim struct {
	userID uuid.UUID
}

func (c Claim) GetUserID() uuid.UUID {
	return c.userID
}
