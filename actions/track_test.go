package actions

// Tests valid URI
func (as *ActionSuite) Test_TrackHandler() {
	res := as.JSON("/track").Get()
	as.Equal(200, res.Code)
}

// Tests fetching valid data from DB
func (as *ActionSuite) Test_TrackGetter() {
	res := as.JSON("/track?tid=12346").Get()
	as.Equal(200, res.Code)
	// as.Contains(res.Body.String(), "Test")
}
