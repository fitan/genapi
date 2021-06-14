package handler

//type SwagUserCallBody struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}
//
//// @Accept  json
//// @Produce  json
//// @Param body body SwagUserCallBody true "这是body\n "
//// @Param query query logic.UserCallQuery false "这是query\n"
//// @Param id path string true "查询id\n多次查询id 的结果\n"
//// @Param name header string false "这是Header 的姓名\n"
//// @Param age header string false "这是age 的姓名\n"
//// @Param ageSub header string false "这是age sub 的姓名\n"
//// @Success 200 {object} Result{data=[]ent.User}
//// @Router /api/usercall [get]
//func UserCall(c *gin.Context) (data interface{}, err error) {
//	in := &logic.UserCallIn{}
//
//	err = c.ShouldBindJSON(&in.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	err = c.ShouldBindUri(&in.Uri)
//	if err != nil {
//		return nil, err
//	}
//
//	err = c.ShouldBindHeader(&in.Header)
//	if err != nil {
//		return nil, err
//	}
//
//	err = c.ShouldBindQuery(&in.Query)
//	if err != nil {
//		return nil, err
//	}
//
//	return logic.UserCall(c, in)
//}
