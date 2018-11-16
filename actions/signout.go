package actions

import "github.com/gobuffalo/buffalo"

// SignoutHandler is a handler to serve up
// a signup page.
func SignoutHandler(c buffalo.Context) error {
	c.Cookies().Delete("user_id")
	//TODO: render login page if cookie not found
	return c.Render(200, r.String("Here will be replaced by login page"))
}
