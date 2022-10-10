package commands

type ConfirmTradeOffer struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
