package model

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.time
	UpdatedAt time.time
}
