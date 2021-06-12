package wrapper

var (
	ErrRedisService = New500Response("redis服务器错误")
	ErrRedisSet     = New400Response("设置redis失败")
	ErrRedisGet     = New400Response("读取redis失败")
)
