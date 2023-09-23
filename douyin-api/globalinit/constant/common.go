package constant

const (
	ErrorMsg   = "网络异常~请稍后"
	SuccessMsg = "ok"
	MySQL_DSN  = "root:0927@tcp("youraddress")/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	Addr       = "youraddress" // Redis 服务器地址和端口
	Password   = "yourpassword"             // Redis 服务器密码，如果没有设置密码则为空字符串
	DB         = 0                    // Redis 数据库索引，默认为 0
	WordFilterFilePath     = "/home/runner/app/pub_sms_banned_words.txt"
	RedisUpdateTime        = 5//秒数
	MaxConcurrency         = 5000
)
