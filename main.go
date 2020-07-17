package main

import (
	"fmt"
	"github.com/felix-xqs/go_example/conf"
	"github.com/felix-xqs/go_example/controller/msgpush"
	"github.com/felix-xqs/golog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	initGin()
	initRouter(conf.Gin)
}
func initRouter(g *gin.Engine) {
	msgpush.InitRouter(g)

	g.Any("/health", Health)
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", conf.C.App.Port)
	err := utils.ListenAndServe(addr, conf.Gin)

	if err != nil {
		golog.Error("err: %v, addr: %v", err, addr)
	}
}
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, "SUCCESS")
}
func initGin() {
	conf.Gin = gin.New()
	conf.Gin.Use(gin.Recovery())
	//conf.Gin.Use(middleware.Boss())
}
