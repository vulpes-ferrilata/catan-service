package commands

type BuildSettlementAndRoad struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
	LandID string `validate:"required,objectid"`
	PathID string `validate:"required,objectid"`
}
