package controller

import (
	"articleManager/logic"
	"articleManager/wxutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//登录
func LoginIn(ctx *gin.Context){
	status:=checkLoginStatus(ctx)
	//判断是否已经登录
	if status==true {
		wxutil.ResponseData(ctx,"login",nil,http.StatusOK)
		return
	}
	//登录
	password:=ctx.Request.FormValue(wxutil.C_PassWordName)
	access:=logic.LoginAccess(password)
	//通过，种cookiet
	if access==true {
		//设置会话可用
		ctx.Set(wxutil.C_LoginStatus,"live")
		//设置全局变量
		genId,err:=wxutil.GetGeneralId()
		if err!=nil {
			log.Printf("LoginIn GetGeneralId error,err:%+v\n",err)
			wxutil.ResponseData(ctx,"",err,0)
			return
		}
		//种cook
		cook:=&http.Cookie{
			Name:       wxutil.C_CookietName,
			Value:      strconv.Itoa(int(genId)),
			Path:       "/",
			MaxAge:     config.CookietLong,
		}
		http.SetCookie(ctx.Writer,cook)
		//写redis
		logic.SetUserTime(strconv.Itoa(int(genId)),"1",config.LoginStatusLong)
		wxutil.ResponseData(ctx,"login",nil,http.StatusOK)
		return
	}
	wxutil.ResponseData(ctx,"",fmt.Errorf("unlogin"),0)
}
