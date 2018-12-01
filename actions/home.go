package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	trackOne := "Web Developer Track"
	trackTwo := "Graphic Designer Track"
	trackThree := "Data Expert Software Track"
	trackOneDes := "Grow as a web developer with us!"
	trackTwoDes := "Equip yourself with graphic design skills!"
	trackThreeDes := "Become a data analyzing expert!"
	courseOne := "Excel 1"
	courseTwo := "Adobe Illustrator 1"
	courseThree := "JavaScript"
	courseOneDes := "An introduction to Excel."
	courseTwoDes := "A first look at Adobe Illustrator."
	courseThreeDes := "A beginner's guide to Javascript."
	return c.Render(200, r.JSON(map[string]string{
		"track_1":      trackOne,
		"track_1_des":  trackOneDes,
		"track_2":      trackTwo,
		"track_2_des":  trackTwoDes,
		"track_3":      trackThree,
		"track_3_des":  trackThreeDes,
		"course_1":     courseOne,
		"course_1_des": courseOneDes,
		"course_2":     courseTwo,
		"course_2_des": courseTwoDes,
		"course_3":     courseThree,
		"course_3_des": courseThreeDes}))
}
