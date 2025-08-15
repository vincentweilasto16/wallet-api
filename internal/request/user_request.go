package request

type GetUserByIDRequest struct {
	UserID string `uri:"id" binding:"required,uuid"`
}
