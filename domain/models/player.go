package models

import (
	"github.com/VulpesFerrilata/catan-service/domain/models/common"
	"github.com/google/uuid"
)

func NewPlayer(id uuid.UUID, userID uuid.UUID) *Player {
	return &Player{
		Entity: common.NewEntity(id),
		userID: userID,
	}
}

type Player struct {
	common.Entity
	userID uuid.UUID
}

func (p Player) GetUserID() uuid.UUID {
	return p.userID
}
