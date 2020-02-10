package wxutil

import "github.com/gin-gonic/gin"

const(
	C_PhotoName="images"
	C_TypeName="gender"
	C_ContextName="content"
)

type ResponseInfo struct {
	Data interface{}
	Message string
	ErrorCode int
}

//返回json消息
func ResponseData(ctx *gin.Context,data interface{},err error,errorCode int){
	responseInfo:=&ResponseInfo{
		Data:      nil,
		Message:   "",
		ErrorCode: 0,
	}
	//初始化返回信息
	if err==nil {
		responseInfo.Data=data
		responseInfo.ErrorCode=errorCode
	}else{
		responseInfo.Message=err.Error()
		responseInfo.ErrorCode=errorCode
	}
	//返回json
	ctx.JSON(errorCode,responseInfo)
}

//得到错误码

func GetErrorCode(err string)int{
	return 0
}