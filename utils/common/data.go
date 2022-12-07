package common

// JsonResult
// 返回数据格式Json
//
type JsonResult struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Count string      `json:"count "`
	Msg   string      `json:"msg "`
}
