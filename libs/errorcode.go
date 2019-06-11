package libs

var codes map[int]string = map[int]string{
	0:     "success",
	-999:  "服务器出了点问题, 请再试一次",
	-1000: "用户已存在",
	-1002: "用户不存在",
	-1003: "密码错误",
	-1004: "无效的cookie",
	-1005: "无效的token",
	-1006: "token已过期",
}
