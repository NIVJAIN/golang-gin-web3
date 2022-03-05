package forms

//EmailStruct ...
type EmailStruct struct {
	EmailID string `json:"email" binding:"required"`
}
