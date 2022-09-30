package commands

type PlayRoadBuildingCard struct {
	GameID  string   `validate:"required,objectid"`
	UserID  string   `validate:"required,objectid"`
	PathIDs []string `validate:"required,min=1,max=2,unique"`
}
