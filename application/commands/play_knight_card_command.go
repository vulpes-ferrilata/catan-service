package commands

type PlayKnightCardCommand struct {
	UserID    string `validate:"required,objectid"`
	GameID    string `validate:"required,objectid"`
	TerrainID string `validate:"required,objectid"`
	PlayerID  string
}
