package profile

import "time"

type Profile struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Biography string    `json:"biography"`
}
