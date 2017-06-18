package kuaidi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	//user 默认用户名配置
	user = "test"
	//password 默认密码配置
	password = "12345678"
	//URL 默认url地址
	URL = "http://capi.yuntongzhi.vip/Api/Sms/sendSms"
	//SmsType 短信验证码
	smsType = int(1)
	//VoiceType 语音验证码
	voiceType = int(2)
	//typeErr 验证码类型错误提示
	typeErr = "不支持的发送类型"
	//mobileErr 手机号码错误提示
	mobileErr = "手机号码格式错误"
	//codeLenErr 验证码长度错误
	codeLenErr = "验证码长度不符"
	//Template
	template = "【云通知】您的验证码是：【变量】。请不要把验证码泄露给其他人。如非本人操作，可不用理会"
	//模板变量
	variable = "【变量】"
)

//Auth 认证信息
type auth struct {
	//User 用户名
	user string
	//Password 密码
	password string
}

//Notice 通知结构
type Notice struct {
	auth     *auth
	URL      string
	Template string
	Variable string
}

//Response 借口返回信息
type Response struct {
	Code       int    `json:"code"`
	Info       string `json:"info"`
	ReturnCode int    `json:"return_code"`
	ReturnInfo string `json:"return_info"`
}

//Send 发送验证码消息 code 需要发送的验证码，sendType 1发送短信验证码 2 发送语音验证码
func (n *Notice) Send(mobile string, code string, sendType int) error {
	if sendType != voiceType && sendType != smsType {
		return errors.New(typeErr)
	}
	if !n.isMobile(mobile) {
		return errors.New(mobileErr)
	}
	codeLen := len(code)
	if codeLen < 4 || codeLen > 8 {
		return errors.New(codeLenErr)
	}
	formData := make(url.Values)
	timeNow := time.Now().Unix()
	formData.Set("content", n.getContent(code, sendType, template))
	formData.Set("times", strconv.FormatInt(timeNow, 10))
	formData.Set("user", n.auth.user)
	formData.Set("passwd", n.encryption(timeNow))
	formData.Set("stype", strconv.Itoa(sendType))
	formData.Set("mobile", mobile)
	var response *http.Response
	var err error
	var body []byte
	response, err = http.PostForm(n.URL, formData)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	Response := new(Response)
	err = json.Unmarshal(body, Response)
	if err != nil {
		return err
	}
	if Response.Code != 0 {
		return errors.New(Response.Info)
	}
	if Response.ReturnCode != 0 {
		return errors.New(Response.ReturnInfo)
	}
	return nil
}

//getCount 组合验证码内容
func (n *Notice) getContent(code string, sendType int, template string) string {
	if sendType == voiceType {
		return code
	}
	return strings.Replace(n.Template, n.Variable, code, -1)
}

//密码加密
func (n *Notice) encryption(times int64) string {
	m := md5.New()
	io.WriteString(m, n.auth.password)
	v := hex.EncodeToString(m.Sum(nil))
	m.Reset()
	io.WriteString(m, v+strconv.FormatInt(times, 10))
	return hex.EncodeToString(m.Sum(nil))
}

//isMobile 是否是手机号码
func (n *Notice) isMobile(mobile string) bool {
	reg := regexp.MustCompile("^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\\\d{8}$")
	return reg.MatchString(mobile)
}

//NewNotice 初始化授权结构
func NewNotice(user string, password string, url string, template string, variable string) *Notice {
	return &Notice{&auth{user, password}, url, template, variable}
}

//SendSms 发送短信验证码,使用默认配置
func SendSms(mobile string, code string) error {
	n := &Notice{&auth{user, password}, URL, template, variable}
	return n.Send(mobile, code, smsType)
}

//SendVoice 发送语音验证码，使用默认配置
func SendVoice(mobile string, code string) error {
	n := &Notice{&auth{user, password}, URL, template, variable}
	return n.Send(mobile, code, voiceType)
}

//RandInt 生成区间随机数
func RandInt(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max-min) + min
}
