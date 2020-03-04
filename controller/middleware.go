package controller

import (
	"articleManager/logic"
	"articleManager/wxutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"log"
)

//登录中间件
func LoginMiddleWare(contextG *gin.Context){
	if checkLoginStatus(contextG)==false {
		wxutil.ResponseData(contextG,"",fmt.Errorf("unlogin"),0)
		contextG.Abort()
		return
	}
	contextG.Next()
}

//判断是否需要登录
func checkLoginStatus(ctx *gin.Context)bool{
	//判断session
	v,exist:=ctx.Get(wxutil.C_LoginStatus)
	if exist==true&&v=="live" {
		return true
	}
	//获得cookiet
	cook,err:=ctx.Request.Cookie(wxutil.C_CookietName)
	if err!=nil {
		return false
	}
	return logic.UserIsLogin(cook.Value,config.LoginStatusLong)
}

//https中间件
func HttpSHandler()gin.HandlerFunc {
	return func(context *gin.Context) {
		secureProcess:=secure.New(secure.Options{
			SSLRedirect:                     true,
			SSLHost:                         "localhost:8089",
		})
		err:=secureProcess.Process(context.Writer,context.Request)
		if err!=nil {
			log.Printf("HttpSHandler Process error,err:%+v\n",err)
			return
		}
		context.Next()
	}
}
