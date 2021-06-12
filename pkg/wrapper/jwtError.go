package wrapper

var (
	ErrJwtAuthToken = New400Response("token生成失败")
	ErrJwtExp       = New400Response("token过期")
	ErrJwtBad       = New400Response("token无效")
	ErrJwtEmpty     = New400Response("token为空")
)
