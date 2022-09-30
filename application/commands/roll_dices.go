package commands

type RollDices struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
