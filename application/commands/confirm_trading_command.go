package commands

type ConfirmTradingCommand struct {
	UserID string `validate:"required,objectid"`
	GameID string `validate:"required,objectid"`
}
