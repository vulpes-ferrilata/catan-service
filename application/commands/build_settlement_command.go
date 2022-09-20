package commands

type BuildSettlementCommand struct {
	UserID string `validate:"required,objectid"`
	GameID string `validate:"required,objectid"`
	LandID string `validate:"required,objectid"`
}
