package models

// course_track_map model is a struct that contains the table id, course id and track id.
type courseTrackMap struct {
	ID       int `db:"id"`
	CourseID int `db:"course_id"`
	TrackID  int `db:"track_id"`
}
