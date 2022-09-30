package commands

type MaritimeTrade struct {
	GameID           string `validate:"required,objectid"`
	UserID           string `validate:"required,objectid"`
	ResourceCardType string `validate:"required"`
}
