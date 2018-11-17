package actions

func (as *ActionSuite) Test_EnrollGetHandler() {
	res := as.HTML("/enroll").Get()
	as.Equal(200, res.Code)
}

// func (as *ActionSuite) Test_EnrollPostHandler() {
// 	res := as.HTML("/enroll").Post("?uid=12345&cid=123")
// 	as.Equal(200, res.Code)
// }
