package actions

import "github.com/gobuffalo/buffalo"

// EnrollGetHandler is a handler to manage course enrollment
func EnrollGetHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "This is the enroll page!" +
		"- you typed " + c.Param("uid") + " for your uid"}))
}

// EnrollPostHandler is a handler to manage course enrollment
func EnrollPostHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "This is the enroll page!" +
		"- you typed " + c.Param("uid") + " for your uid " +
		c.Param("cid") + " for your cid"}))
}
