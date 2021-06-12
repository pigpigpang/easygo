package wrapper

var (
	ErrEmptySendUser     = New400Response("发件人不存在")
	ErrEmptyRcptUser     = New400Response("收件人不存在")
	ErrInsertLetter      = New400Response("发信失败")
	ErrInvalidLetterType = New400Response("信件类型错误")
	ErrLetterDelete      = New400Response("信件删除失败")
	ErrFileUpload        = New400Response("附件上传失败")
)
