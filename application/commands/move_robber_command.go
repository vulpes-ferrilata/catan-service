package commands

type MoveRobberCommand struct {
	UserID    string `validate:"required,objectid"`
	GameID    string `validate:"required,objectid"`
	TerrainID string `validate:"required,objectid"`
	PlayerID  string
}
