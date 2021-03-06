package actions

// Tests valid URI
func (as *ActionSuite) Test_CourseHandler() {
	res := as.JSON("/courses").Get()
	as.Equal(200, res.Code)
}

// Tests fetching valid data from DB
func (as *ActionSuite) Test_CourseGetter() {
	res := as.JSON("/courses?cid=12346").Get()
	as.Equal(200, res.Code)
	// as.Contains(res.Body.String(), "Test")
}
