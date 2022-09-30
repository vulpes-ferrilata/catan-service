package commands

type CancelTrading struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
