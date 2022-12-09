package common

// JsonResult
// 返回数据格式Json
//
/*type JsonResult struct {
	Code  int         `json:"code" form:"code"`
	Data  interface{} `json:"data" form:"data"`
	Count int      `json:"count " form:"count"`
	Msg   string      `json:"msg " form:"msg"`
}*/
type JsonResult struct {
	Code  int         `json:"code" `
	Data  interface{} `json:"data" `
	Count int         `json:"count " `
	Msg   string      `json:"msg "`
}
