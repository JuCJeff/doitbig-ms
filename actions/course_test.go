package actions

func (as *ActionSuite) Test_CourseHandler() {
	res := as.JSON("/course").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "This is the course page!")
}

// func (as *ActionSuite) Test_TrackGetter() {
// 	res := as.JSON("/track?uid=12345").Get()
// 	as.Equal(200, res.Code)
// 	as.Contains(res.Body.String(), "First Name: John")
// }
