package commands

type BuyDevelopmentCard struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
