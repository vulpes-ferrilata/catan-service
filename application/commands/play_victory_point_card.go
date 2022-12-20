package commands

type PlayVictoryPointCard struct {
	GameID            string `validate:"required,objectid"`
	UserID            string `validate:"required,objectid"`
	DevelopmentCardID string `validate:"required,objectid"`
}
