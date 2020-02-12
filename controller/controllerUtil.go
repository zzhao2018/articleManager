package controller

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/smtp"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"articleManager/conf"
	"articleManager/logic"
)



/*******************解析配置文件********************/
var config *conf.Conf

//设置全局配置文件
func SetConfig(config_in *conf.Conf){
	config=config_in
}



/******************定时任务******************/
//默认类别、时间
var(
	//配置信息
	defaultUserType=[]int{}
	defaultHour=[]int{}
	defaultMin=[]int{}
	//上下文设置
	defaultcCancleFunc context.CancelFunc
	defauleCtx context.Context
	//锁
	sendLock sync.Mutex
)


//重置默认时间
func SetDefaultSendParam(useType []int,hours []int,mins []int){
	defaultUserType=useType
	defaultHour=hours
	defaultMin=mins
	//设置cancel函数
	ctx,cancle:=context.WithCancel(context.TODO())
	defauleCtx=ctx
	defaultcCancleFunc=cancle
}


/**************************定时任务工作类************************/
type TickJobStruct struct {
	typeI int
}

func NewTickJob(typeI int) *TickJobStruct {
	return &TickJobStruct{typeI: typeI}
}

//任务动作
func (t *TickJobStruct) Run() {
	//随机获取文章
	articleInfo:=logic.SearchRandomArticle(t.typeI)
	if articleInfo.Article_context=="" {
		log.Printf("type:%s has no message!",t.typeI)
		return
	}
	typeInfo:=logic.SearchTypeById(int(articleInfo.Type))
	//解析图片
	var photoArr []string
	err:=json.Unmarshal([]byte(articleInfo.Photo),&photoArr)
	if err!=nil {
		log.Printf("Run json unmarshal error,photo:%+v,err:%+v\n",articleInfo.Photo,err)
		return
	}
	//发送
	DoJob(typeInfo.TypeName,articleInfo.Article_context,photoArr)
}



/*******************发送邮件*******************/
type MailMessage struct {
	typeS string
	host string
	port int
	password string
	header map[string]string
	message *bytes.Buffer
	email string
	boundary string
	attaFile []string
}

//设置邮件信息
func(m *MailMessage)setMailMess(email string ,toEmail string,subject string,typeS string,body string,photoPath []string)(error){
	m.typeS=typeS
	m.host="smtp.qq.com"
	m.port=465
	m.password="okdhlrvbmyotbiad"
	m.email=email
	m.boundary="ds13difsknfsifuere134"
	m.attaFile=photoPath
	m.header= make(map[string]string)
	m.header["From"] = typeS + "<" + email + ">"
	m.header["To"] = toEmail
	m.header["Subject"] = typeS+"_"+subject
	m.header["Content-Type"] = fmt.Sprintf("multipart/mixed; charset=UTF-8;boundary=%s",m.boundary)
	//设置头部信息
	m.message=bytes.NewBuffer(nil)
	for k, v := range m.header {
		m.message.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	//设置文本头部信息
	m.message.WriteString(fmt.Sprintf("--%s\r\n", m.boundary))
	m.message.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	m.message.WriteString("\r\n" + body)
	for i:=0;i<len(m.attaFile);i++{
		//设置图片头部信息
		m.message.WriteString(fmt.Sprintf("\n--%s\r\n", m.boundary))
		m.message.WriteString("Content-Type: application/octet-stream\r\n")
		m.message.WriteString("Content-Transfer-Encoding: base64\r\n")
		filePathArr:=strings.Split(m.attaFile[i],string(filepath.Separator))
		m.message.WriteString("Content-Disposition: attachment; filename=\"" + typeS+"_"+filePathArr[len(filePathArr)-1] + "\"\r\n\r\n")
		//读取附件
		attaData,err:=ioutil.ReadFile(m.attaFile[i])
		if err!=nil {
			log.Printf("setMailMess read photo error,err:%+v\n",err)
			return err
		}
		byteData:=make([]byte,base64.StdEncoding.EncodedLen(len(attaData)))
		base64.StdEncoding.Encode(byteData,attaData)
		m.message.Write(byteData)
	}
	return nil
}


//发送邮件至一人
func doSend(message *MailMessage){
	auth := smtp.PlainAuth(
		"",
		message.email,
		message.password,
		message.host,
	)
	toEmail:=message.header["To"]
	//发送邮件
	err := SendMailUsingTLS(
		fmt.Sprintf("%s:%d", message.host, message.port),
		auth,
		message.email,
		[]string{toEmail},
		message.message.Bytes(),
	)
	if err != nil {
		log.Printf("Dialing Error:%+v\n", err)
		return
	}
}

//发送邮件任务至多人
func DoJob(typeS string,body string,photoPath []string) {
	for i:=0;i<len(config.ToEmailAddr);i++  {
		message:=&MailMessage{}
		message.setMailMess(config.EmailAddr,config.ToEmailAddr[i],config.Subject,typeS,body,photoPath)
		log.Println("***************************")
		log.Println("begin send email:",time.Now())
		doSend(message)
		fmt.Println("send success:",time.Now())
		fmt.Println("***************************")
	}
}


/***************发送邮件函数**************/
//返回smtp客户端
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}


//参考net/smtp的func SendMail()
func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {
	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
