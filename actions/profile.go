package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"gitlab.com/doitbig-ms/models"
)

// ProfileHandler is a handler to serve up
// a signup page.
func ProfileHandler(c buffalo.Context) error {
	student := models.Student{}
	db, _ := pop.Connect("development")
	db.Find(&student, c.Param("uid"))
	return c.Render(200, r.JSON(student))
}
