package kuaidi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	//user 默认用户名配置
	user = "test"
	//password 默认密码配置
	password = "123456"
	//URL 默认url地址
	URL = "http://capi.yuntongzhi.vip/Api/Sms/sendSms"
	//SmsType 短信验证码
	smsType = int(1)
	//VoiceType 语音验证码
	voiceType = int(2)
	//typeErr 验证码类型错误提示
	typeErr = "不支持的发送类型"
)

//Auth 认证信息
type Auth struct {
	//User 用户名
	user string
	//Password 密码
	password string
}

//Notice 通知结构
type Notice struct {
	auth *Auth
	URL  string
}

//Response 借口返回信息
type Response struct {
	Code       int    `json:"code"`
	Info       string `json:"info"`
	ReturnCode int    `json:"return_code"`
	ReturnInfo string `json:"return_info"`
}

//Send 发送验证码消息 code 需要发送的验证码，sendType 1发送短信验证码 2 发送语音验证码
func (n *Notice) Send(mobile string, code int64, sendType int) error {
	if sendType != voiceType && sendType != smsType {
		return errors.New(typeErr)
	}
	formData := make(url.Values)
	timeNow := time.Now().Unix()
	formData.Set("content", strconv.FormatInt(code, 10))
	formData.Set("times", strconv.FormatInt(timeNow, 10))
	formData.Set("user", n.auth.user)
	formData.Set("passwd", n.encryption(timeNow))
	formData.Set("stype", strconv.Itoa(sendType))
	fmt.Println(formData)
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
	fmt.Println(string(body))
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

//密码加密
func (n *Notice) encryption(times int64) string {
	m := md5.New()
	io.WriteString(m, string(n.auth.password))
	v := m.Sum(nil)
	passwordMd5 := hex.EncodeToString(v) + strconv.FormatInt(times, 10)
	io.WriteString(m, string(passwordMd5))
	return hex.EncodeToString(m.Sum(nil))
}

//NewNotice 初始化授权结构
func NewNotice(user string, password string, url string) *Notice {
	return &Notice{&Auth{user, password}, url}
}

//SendSms 发送短信验证码,使用默认配置
func SendSms(mobile string, code int64) error {
	n := &Notice{&Auth{user, password}, URL}
	return n.Send(mobile, code, smsType)
}

//SendVoice 发送语音验证码，使用默认配置
func SendVoice(mobile string, code int64) error {
	n := &Notice{&Auth{user, password}, URL}
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
