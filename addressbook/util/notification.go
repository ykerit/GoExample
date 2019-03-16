package util

// 接口：
// 拨打电话通知
// 增删改查通知
// 呼叫提醒
type notificationer interface {
	notifiy()
	callReminder()
}

