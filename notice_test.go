package notice

import "testing"

const mobile = "18025434220"

//Test_SendByAuth 测试指定授权信息的发送
func Test_SendByAuth(t *testing.T) {
	notice := NewNotice("test", "123456", "http://capi.yuntongzhi.vip/Api/Sms/sendSms")
	err := notice.Send(mobile, int64(1234), 1)
	if err != nil {
		t.Error(err)
	}
}

//Test_Send 测试使用默认授权信息的发送
func Test_Send(t *testing.T) {
	err := SendSms(mobile, int64(444))
	if err != nil {
		t.Error(err)
	}
	err = SendVoice(mobile, int64(333))
	if err != nil {
		t.Error(err)
	}
}

//Test_RandInt 测试随机数生成
func Test_RandInt(t *testing.T) {
	for i := 0; i < 2; i++ {
		// fmt.Println(RandInt(1000, 9999))
	}
}
