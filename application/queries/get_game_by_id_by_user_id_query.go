package queries

type GetGameByIDByUserIDQuery struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
