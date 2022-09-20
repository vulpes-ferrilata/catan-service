package commands

type BuyDevelopmentCardCommand struct {
	UserID string `validate:"required,objectid"`
	GameID string `validate:"required,objectid"`
}
