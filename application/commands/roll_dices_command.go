package commands

type RollDicesCommand struct {
	UserID string `validate:"required,objectid"`
	GameID string `validate:"required,objectid"`
}
