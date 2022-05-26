package commands

type CreateGameCommand struct {
	UserID string `validate:"required,uuid4"`
	GameID string `validate:"required,uuid4"`
}
