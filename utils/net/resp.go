package net

import (
	"github.com/gin-gonic/gin"
	"github.com/xqs-go-lib/gosdk/lib"
	"net/http"
)

//这里和gosdk/mygin 里的区别就是，那里是用的自己封装了一层context，这里直接用gin的context就可以

func Reply200(ctx *gin.Context, resp *lib.Resp) {
	ctx.JSON(http.StatusOK, resp)
}

func FailWithDetail(code lib.ErrCode, detail string) *lib.Resp {
	return &lib.Resp{
		Code:   code,
		Msg:    lib.CodeMap[code],
		Detail: detail,
	}
}

func FailWithMsg(code lib.ErrCode, msg string, detail string) *lib.Resp {
	return &lib.Resp{
		Code:   code,
		Msg:    msg,
		Detail: detail,
	}
}

func OkWithData(data interface{}) *lib.Resp {
	return &lib.Resp{
		Code: lib.CodeOk,
		Data: data,
	}
}

func Ok() *lib.Resp {
	return &lib.Resp{
		Code: lib.CodeOk,
	}
}

// Data ...
func Data(data interface{}) *lib.Resp {
	return &lib.Resp{
		Data: data,
	}
}
