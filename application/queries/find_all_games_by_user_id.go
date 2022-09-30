package queries

type FindGamesByUserID struct {
	UserID string `validate:"required,objectid"`
}
