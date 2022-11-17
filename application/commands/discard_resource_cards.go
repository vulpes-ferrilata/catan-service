package commands

type DiscardResourceCards struct {
	GameID          string   `validate:"required,objectid"`
	UserID          string   `validate:"required,objectid"`
	ResourceCardIDs []string `validate:"required,unique"`
}
