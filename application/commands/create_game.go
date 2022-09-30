package commands

type CreateGame struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
