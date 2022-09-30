package commands

type ConfirmTrading struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
