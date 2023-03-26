package wechat

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}

func MessageHandler(msg *openwechat.Message) {
	if msg.IsText() {
		msg.ReplyText("pong")
		sender, _ := msg.Sender()
		fmt.Printf("nickname:%s, username:%s, ID:%s, isGroupMsg:%v, isAt:%v, fromUsername:%s, toUsername:%s, content:%s\n",
			sender.NickName, sender.UserName, sender.ID(), msg.IsSendByGroup(), msg.IsAt(),
			msg.FromUserName, msg.ToUserName, msg.Content)
	}

	// 单聊的文本消息
	if msg.IsText() && !msg.IsSendByGroup() {

	}
}

func Run() error {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 注册消息处理函数
	bot.MessageHandler = MessageHandler
	// 注册登陆二维码回调
	bot.UUIDCallback = ConsoleQrCode

	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	// 登陆
	if err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		return err
	}
	_ = bot.Block()
	return nil
}
