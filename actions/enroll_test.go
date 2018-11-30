package actions

func (as *ActionSuite) Test_EnrollGetHandler() {
	res := as.HTML("/enroll").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_EnrollPostHandler() {
	res := as.HTML("/enroll").Post("?uid=12346&cid=123")
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_EnrollPostHandlerErr() {
	res := as.HTML("/enroll").Post("?uid=test&cid=test")
	as.Equal(500, res.Code)
}
