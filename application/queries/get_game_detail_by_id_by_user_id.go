package queries

type GetGameDetailByIDByUserID struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}
