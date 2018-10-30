package actions

import "github.com/gobuffalo/buffalo"

// SignoutHandler is a handler to serve up
// a signup page.
func SignoutHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "This is the signout page!"}))
}
