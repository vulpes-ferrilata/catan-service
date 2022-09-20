package commands

type OfferTradingCommand struct {
	UserID   string `validate:"required,objectid"`
	GameID   string `validate:"required,objectid"`
	PlayerID string `validate:"required,objectid"`
}
