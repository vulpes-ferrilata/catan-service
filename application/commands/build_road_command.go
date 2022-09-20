package commands

type BuildRoadCommand struct {
	UserID string `validate:"required,objectid"`
	GameID string `validate:"required,objectid"`
	PathID string `validate:"required,objectid"`
}
