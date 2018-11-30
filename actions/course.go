package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"gitlab.com/doitbig-ms/models"
)

// CourseHandler is a handler to serve up
// a course page.
func CourseHandler(c buffalo.Context) error {
	course := models.Course{}
	db, _ := pop.Connect("development")
	db.Find(&course, c.Param("cid"))
	return c.Render(200, r.JSON(course))
}
