package msgpush

import "github.com/gin-gonic/gin"
//路由文件，只定义路由
func InitRouter(app *gin.Engine){
	group := app.Group("/v1")
	group.POST("notify/pay/success",PaySuccessNotice)
}