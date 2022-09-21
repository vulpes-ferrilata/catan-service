package commands

type PlayRoadBuildingCardCommand struct {
	UserID  string   `validate:"required,objectid"`
	GameID  string   `validate:"required,objectid"`
	PathIDs []string `validate:"required,min=1,max=2,unique"`
}
