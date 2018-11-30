package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	trackOne := "Web Developer Track"
	trackTwo := "STS Track Two"
	trackThree := "STS Track Three"
	trackOneDes := "STS Track One description goes here."
	trackTwoDes := "STS Track Two description goes here."
	trackThreeDes := "STS Track Three description goes here."
	courseOne := "STS Course One"
	courseTwo := "STS Course Two"
	courseThree := "STS Course Three"
	courseOneDes := "STS Course One Description goes here."
	courseTwoDes := "STS Course Two Description goes here."
	courseThreeDes := "STS Course Three Description goes here."
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
