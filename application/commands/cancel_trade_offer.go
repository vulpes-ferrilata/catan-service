package commands

type CancelTradeOffer struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
