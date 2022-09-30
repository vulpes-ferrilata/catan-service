package commands

type UpgradeCity struct {
	GameID         string `validate:"required,objectid"`
	UserID         string `validate:"required,objectid"`
	ConstructionID string `validate:"required,objectid"`
}
