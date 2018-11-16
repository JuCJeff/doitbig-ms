package models

// Time model is a struct that contains the id, year, month, day, startTime, and endTime of time struct.
type Time struct {
	ID        int    `db:"id"`
	Year      string `db:"year"`
	Month     string `db:"month"`
	Day       string `db:"day"`
	StartTime string `db:"start_time"`
	EndTime   string `db:"end_time"`
}
