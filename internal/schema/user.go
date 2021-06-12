package schema

type RegisterReqBodySchema struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Personal string `json:"personal"`
}
type RegisterResBodySchema struct {
	Msg string
}

type LoginReqBodySchema struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResBodySchema struct {
	Msg      string
	Personal string
	Token    string
}
