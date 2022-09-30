package commands

type BuildRoad struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
	PathID string `validate:"required,objectid"`
}
