package commands

type PlayMonopolyCardCommand struct {
	UserID           string `validate:"required,objectid"`
	GameID           string `validate:"required,objectid"`
	ResourceCardType string `validate:"required"`
}
