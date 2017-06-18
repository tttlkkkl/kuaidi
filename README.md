# kuaidi
云通知go语言包
 
安装：
go get github.com/tttlkkkl/kuaidi

 
使用示例：
```golang
package main

import (
	"fmt"
	"strconv"

	"github.com/tttlkkkl/kuaidi"
)

func main() {
	mobile := "18000000000"
	notice := kuaidi.NewNotice("test", "12345678", "http://capi.yuntongzhi.vip/Api/Sms/sendSms", "【云通知】您的验证码是：【变量】。请不要把验证码泄露给其他人。如非本人操作，可不用理会", "【变量】")
	//发送短信验证码
	err := notice.Send(mobile, strconv.FormatInt(kuaidi.RandInt(1000, 9999), 10), 1)
	if err != nil {
		fmt.Println(err)
	}
	//发送语音验证码
	err = notice.Send(mobile, strconv.FormatInt(kuaidi.RandInt(1000, 9999), 10), 2)
	if err != nil {
		fmt.Println(err)
	}
	//发送短信验证码，使用默认配置，可以直接修改源码包中的配置
	err = kuaidi.SendSms(mobile, strconv.FormatInt(kuaidi.RandInt(1000, 9999), 10))
	if err != nil {
		fmt.Println(err)
	}
	//发送语音验证码,使用默认配置，可以直接修改源码包中的配置
	err = kuaidi.SendVoice(mobile, strconv.FormatInt(kuaidi.RandInt(1000, 9999), 10))
	if err != nil {
		fmt.Println(err)
	}
}


```
