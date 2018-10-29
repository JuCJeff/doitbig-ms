package actions

import "github.com/gobuffalo/buffalo"

// SignupHandler is a handler to serve up
// a signup page.
func SignupHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "This is the signup page!"}))
}
