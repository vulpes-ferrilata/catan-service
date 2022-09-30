package commands

type StartGame struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
