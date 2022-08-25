package commands

type BuildSettlementAndRoadCommand struct {
	UserID string `validate:"required,objectid"`
	GameID string `validate:"required,objectid"`
	LandID string `validate:"required,objectid"`
	PathID string `validate:"required,objectid"`
}
