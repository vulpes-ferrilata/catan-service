package entities

import (
	"github.com/VulpesFerrilata/catan-service/persistence/entities/common"
	"github.com/google/uuid"
)

type Achievement struct {
	common.Entity
	GameID   uuid.UUID
	PlayerID *uuid.UUID
	Type     string
}
