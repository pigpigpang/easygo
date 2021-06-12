package wrapper

// 定义错误
var (
	ErrBadRequest              = New400Response("请求发生错误")
	ErrInvalidParent           = New400Response("无效的父级节点")
	ErrNotAllowDeleteWithChild = New400Response("含有子级，不能删除")
	ErrNotAllowDelete          = New400Response("资源不允许删除")

	ErrNoPerm          = NewResponse(401, 401, "无访问权限")
	ErrInvalidToken    = NewResponse(9999, 401, "令牌失效")
	ErrNotFound        = NewResponse(404, 404, "资源不存在")
	ErrMethodNotAllow  = NewResponse(405, 405, "方法不被允许")
	ErrTooManyRequests = NewResponse(429, 429, "请求过于频繁")
	ErrInternalServer  = NewResponse(500, 500, "服务器发生错误")
)
