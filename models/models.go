package models

import (
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

func init() {
	env := envy.Get("GO_ENV", "development")
	DB, _ = pop.Connect(env)
	pop.Debug = env == "development"
}
