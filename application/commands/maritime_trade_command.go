package commands

type MaritimeTradeCommand struct {
	UserID           string `validate:"required,objectid"`
	GameID           string `validate:"required,objectid"`
	ResourceCardType string `validate:"required"`
}
