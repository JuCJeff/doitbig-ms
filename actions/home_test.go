package actions

func (as *ActionSuite) Test_HomeHandler() {
	res := as.JSON("/").Get()
	as.Equal(200, res.Code)
}
