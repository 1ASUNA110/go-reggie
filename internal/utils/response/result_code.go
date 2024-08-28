package response

type ResultCode struct {
	Code int    // 状态码
	Msg  string // 异常信息
}

func SERVER_ERROR() ResultCode {
	return ResultCode{Code: 0, Msg: "失败"}
}

func SUCCESS() ResultCode {
	return ResultCode{Code: 1, Msg: "成功"}
}

func PARAM_ERROR() ResultCode {
	return ResultCode{Code: 2, Msg: "参数错误"}
}

func USER_LOGIN_ERROR() ResultCode {
	return ResultCode{Code: 3, Msg: "用户名或密码错误"}
}

func User_DISABLED_ERROR() ResultCode {
	return ResultCode{Code: 4, Msg: "账号已禁用"}
}

func LOGIN_CHECK_ERROR() ResultCode { return ResultCode{Code: 5, Msg: "未登录"} }

func USER_IS_EXIST() ResultCode { return ResultCode{Code: 6, Msg: "用户已存在"} }
