package commands

type UpgradeCityCommand struct {
	UserID         string `validate:"required,objectid"`
	GameID         string `validate:"required,objectid"`
	ConstructionID string `validate:"required,objectid"`
}
