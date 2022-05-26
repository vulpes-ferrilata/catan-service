package entities

import (
	"github.com/VulpesFerrilata/catan-service/persistence/entities/common"
	"github.com/google/uuid"
)

type Player struct {
	common.Entity
	GameID uuid.UUID
	UserID uuid.UUID
}
