package commands

type EndTurn struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
