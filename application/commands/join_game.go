package commands

type JoinGame struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
