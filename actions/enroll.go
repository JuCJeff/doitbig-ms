package actions

import "github.com/gobuffalo/buffalo"

// LoginHandler is a handler to serve up
// a signup page.
func EnrollGetHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "This is the enroll page!" + 
			"- you typed " + c.Param("uid") + " for your uid"}))
}