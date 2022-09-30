package commands

type OfferTrading struct {
	GameID   string `validate:"required,objectid"`
	UserID   string `validate:"required,objectid"`
	PlayerID string `validate:"required,objectid"`
}
