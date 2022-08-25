package queries

type FindGamesByUserIDQuery struct {
	UserID string `validate:"required,objectid"`
}
