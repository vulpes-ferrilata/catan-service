package entities

import (
	"github.com/VulpesFerrilata/catan-service/persistence/entities/common"
	"github.com/google/uuid"
)

type Dice struct {
	common.Entity
	GameID uuid.UUID
	Number int
}
