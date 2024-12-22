package entity

type User struct {
	ID       int64  `db:"id" json:"id"`
	Email    string `db:"email" json:"email" form:"email" validate:"required,email,max=255"`
	Password string `db:"password" json:"password" form:"password" validate:"required,max=32"`
}
