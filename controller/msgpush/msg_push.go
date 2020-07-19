package msgpush

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)
//controller 层，只负责前置判断参数合法性，以及提取参数，返回接口请求结果
func PaySuccessNotice(c *gin.Context){
	ctx, err := context.WithTimeout(c,time.Duration(1))
	traceID := ctx.Value("")
	c.Done()
}
