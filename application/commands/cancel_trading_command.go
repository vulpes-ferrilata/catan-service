package commands

type CancelTradingCommand struct {
	UserID string `validate:"required,objectid"`
	GameID string `validate:"required,objectid"`
}
