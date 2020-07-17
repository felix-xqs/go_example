package def

// Code 用于定义返回码
type Code int

// 定义各种返回码
const (
	CodeOk                 Code = 400001
	CodeSrv                     = 400002
	CodePara                    = 400003

)

var CodeMap = map[Code]string{
	CodeSrv:                "服务错误",
	CodePara:               "参数错误",
}

