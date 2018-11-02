package models

// Student model is a struct that contains the id, firstName, lastNAme, email, password, and image of student struct.
type Student struct {
	ID        int    `db:"id"`
	FirstName string `db:"firstname"`
	LastName  string `db:"lastname"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Image     []byte `db:"image"`
}
