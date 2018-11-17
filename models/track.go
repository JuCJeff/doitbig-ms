package models

// Track model is a struct that contains the id and the track name.
type Track struct {
	ID   int    `db:"id"`
	name string `db:"name"`
}
