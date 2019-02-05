package model

//User User
type User struct {
	ID          int    `json:"id" validate:"number" example:"1"`
	MailAddress string `json:"mail_address" validate:"required,email" example:"test@clean-arch-example.com"`
	Timestamps
}

//Users Users
type Users struct {
	User []*User `json:"user"`
}
