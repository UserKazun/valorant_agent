package model

type SuccessResult struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
}

type Agent struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Background string `db:"background" json:"background"`
	CreatedAt  string `db:"created_at" json:"create_at"`
	UpdatedAt  string `db:"updated_at" json:"updated_at"`
}

type User struct {
	ID        int    `db:"id" json:"id"`
	UserName  string `db:"user_name" json:"user_name"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}
