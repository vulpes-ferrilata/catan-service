package commands

type PlayYearOfPlentyCardCommand struct {
	UserID            string   `validate:"required,objectid"`
	GameID            string   `validate:"required,objectid"`
	ResourceCardTypes []string `validate:"required"`
}
