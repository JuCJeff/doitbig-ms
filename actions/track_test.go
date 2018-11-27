package actions

func (as *ActionSuite) Test_TrackHandler() {
	res := as.JSON("/track").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "This is the track page!")
}

// func (as *ActionSuite) Test_TrackGetter() {
// 	res := as.JSON("/track?uid=12345").Get()
// 	as.Equal(200, res.Code)
// 	as.Contains(res.Body.String(), "First Name: John")
// }
