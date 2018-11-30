package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"gitlab.com/doitbig-ms/models"
)

// CourseHandler is a handler to serve up
// a course page.
func CourseHandler(c buffalo.Context) error {
	if c.Param("cid") == "" {
		courses := []models.Course{}
		db := c.Value("tx").(*pop.Connection)
		db.All(&courses)
		return c.Render(200, r.JSON(courses))

	} else {
		course := models.Course{}
		db := c.Value("tx").(*pop.Connection)
		db.Find(&course, c.Param("cid"))
		return c.Render(200, r.JSON(course))
	}
}
