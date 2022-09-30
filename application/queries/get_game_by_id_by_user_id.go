package queries

type GetGameByIDByUserID struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
