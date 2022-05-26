package commands

type JoinGameCommand struct {
	UserID string `validate:"required,uuid4"`
	GameID string `validate:"required,uuid4"`
}
