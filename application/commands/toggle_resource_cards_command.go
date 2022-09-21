package commands

type ToggleResourceCardsCommand struct {
	UserID          string   `validate:"required,objectid"`
	GameID          string   `validate:"required,objectid"`
	ResourceCardIDs []string `validate:"required,unique"`
}
