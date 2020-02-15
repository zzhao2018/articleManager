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
		//种cook
		genId,err:=wxutil.GetGeneralId()
		if err!=nil {
			log.Printf("LoginIn GetGeneralId error,err:%+v\n",err)
			wxutil.ResponseData(ctx,"",err,0)
			return
		}
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