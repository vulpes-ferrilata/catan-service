package commands

type PlayYearOfPlentyCard struct {
	GameID            string   `validate:"required,objectid"`
	UserID            string   `validate:"required,objectid"`
	ResourceCardTypes []string `validate:"required,min=1,max=2"`
}
