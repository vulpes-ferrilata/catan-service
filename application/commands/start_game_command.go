package commands

type StartGameCommand struct {
	UserID string `validate:"required,objectid"`
	GameID string `validate:"required,objectid"`
}
