package kuaidi

import (
	"fmt"
	"strconv"
	"testing"
)

const mobile = "18025434220"

//Test_SendByAuth 测试指定授权信息的发送
func Test_SendByAuth(t *testing.T) {
	notice := NewNotice("test", "12345678", "http://capi.yuntongzhi.vip/Api/Sms/sendSms", "【云通知】您的验证码是：【变量】。请不要把验证码泄露给其他人。如非本人操作，可不用理会", "【变量】")
	err := notice.Send(mobile, strconv.FormatInt(RandInt(1000, 9999), 10), 2)
	if err != nil {
		t.Error(err)
	}
}

//Test_Send 测试使用默认授权信息的发送
func Test_Send(t *testing.T) {
	err := SendSms(mobile, strconv.FormatInt(RandInt(1000, 9999), 10))
	if err != nil {
		t.Error(err)
	}
	err = SendVoice(mobile, strconv.FormatInt(RandInt(1000, 9999), 10))
	if err != nil {
		t.Error(err)
	}
}

//Test_RandInt 测试随机数生成
func Test_RandInt(t *testing.T) {
	for i := 0; i < 2; i++ {
		fmt.Println(strconv.FormatInt(RandInt(1000, 9999), 10))
	}
}
