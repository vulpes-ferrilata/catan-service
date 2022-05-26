package entities

import (
	"github.com/VulpesFerrilata/catan-service/persistence/entities/common"
	"github.com/google/uuid"
)

type Game struct {
	common.Entity
	Status         string
	ActivePlayerID uuid.UUID
	Turn           int
}
