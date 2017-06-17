# kuaidi
云通知go语言包
 
安装：
go get github.com/tttlkkkl/kuaidi

 
使用示例：
```golang
package main

import (
	"fmt"

	"github.com/tttlkkkl/kuaidi"
)

func main() {
	mobile := "18000000000"
	notice := kuaidi.NewNotice("test", "123456", "http://capi.yuntongzhi.vip/Api/Sms/sendSms")
	//发送短信验证码
	err := notice.Send(mobile, kuaidi.RandInt(1000, 9999), 1)
	if err != nil {
		fmt.Println(err)
	}
	//发送语音验证码
	err = notice.Send(mobile, kuaidi.RandInt(1000, 9999), 2)
	if err != nil {
		fmt.Println(err)
	}
	//发送短信验证码，使用默认配置，可以直接修改源码包中的配置
	err := kuaidi.SendSms(mobile, kuaidi.RandInt(1000, 9999))
	if err != nil {
		fmt.Println(err)
	}
	//发送语音验证码,使用默认配置，可以直接修改源码包中的配置
	err = kuaidi.SendVoice(mobile, kuaidi.RandInt(1000, 9999))
	if err != nil {
		fmt.Println(err)
	}
}

```
