package commands

type PlayYearOfPlentyCard struct {
	GameID                     string   `validate:"required,objectid"`
	UserID                     string   `validate:"required,objectid"`
	DevelopmentCardID          string   `validate:"required,objectid"`
	DemandingResourceCardTypes []string `validate:"required,min=1,max=2"`
}
