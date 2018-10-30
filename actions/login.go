package actions

import "github.com/gobuffalo/buffalo"

// LoginHandler is a handler to serve up
// a signup page.
func LoginHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "This is the Login page!"}))
}

