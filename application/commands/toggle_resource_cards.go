package commands

type ToggleResourceCards struct {
	GameID          string   `validate:"required,objectid"`
	UserID          string   `validate:"required,objectid"`
	ResourceCardIDs []string `validate:"required,unique"`
}
