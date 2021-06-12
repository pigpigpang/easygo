package wrapper

var (
	ErrEmptyUserName     = New400Response("用户名为空")
	ErrInvalidUserName   = New400Response("无效的用户名")
	ErrInvalidPassword   = New400Response("无效的密码")
	ErrInvalidUser       = New400Response("无效的用户")
	ErrUserDisable       = New400Response("用户被禁用，请联系管理员")
	ErrInsertUser        = New400Response("注册失败")
	ErrMinPassword       = New400Response("用户密码过短")
	ErrMaxPassword       = New400Response("用户密码过长")
	ErrLevelPassword     = New400Response("密码强度过低，请使用大小写字母和数字组合")
	ErrExistedUser       = New400Response("用户已存在")
	ErrMinUsername       = New400Response("用户名过短")
	ErrMaxUsername       = New400Response("用户名过长")
	ErrBannedUser        = New400Response("用户被封禁")
	ErrInvalidBannedUser = New400Response("封禁用户失败")
)
