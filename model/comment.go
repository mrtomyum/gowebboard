package model

type Comment struct {
	ID        int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
