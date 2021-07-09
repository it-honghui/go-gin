package domain

var (
	SUCCESS   = Recode{Code: 0, Msg: "ok"}
	NOT_FOUND = Recode{Code: 404, Msg: "not found"}
	ERROR     = Recode{Code: 500, Msg: "error"}

	USERNAME_OR_PASSWORD_IS_EMPTY = Recode{Code: 1001, Msg: "用户名或密码为空"}
	USERNAME_ALREADY_EXIST        = Recode{Code: 1002, Msg: "用户名重复"}
	WRONG_PASSWORD                = Recode{Code: 1003, Msg: "密码不正确"}

	TOKEN_NOT_EXIST   = Recode{Code: 1004, Msg: "Token不存在,请重新登陆"}
	TOKEN_EXPIRE      = Recode{Code: 1005, Msg: "Token已过期,请重新登陆"}
	TOKEN_ERROR       = Recode{Code: 1006, Msg: "Token不正确,请重新登陆"}
	TOKEN_TYPE_ERROR  = Recode{Code: 1007, Msg: "Token格式错误,请重新登陆"}
	USER_UNAUTHORIZED = Recode{Code: 1008, Msg: "用户未经授权"}
)

type Recode struct {
	Code uint
	Msg  string
}
