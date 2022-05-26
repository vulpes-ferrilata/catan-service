package commands

type StartGameCommand struct {
	UserID string `validate:"required,uuid4"`
	GameID string `validate:"required,uuid4"`
}
