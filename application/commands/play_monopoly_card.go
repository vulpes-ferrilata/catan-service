package commands

type PlayMonopolyCard struct {
	GameID                    string `validate:"required,objectid"`
	UserID                    string `validate:"required,objectid"`
	DevelopmentCardID         string `validate:"required,objectid"`
	DemandingResourceCardType string `validate:"required"`
}
