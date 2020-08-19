package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	UNAUTHORIZED:   "没有访问权限",

	ERROR_EDIT_USER_FAIL:              "编辑用户失败",
	ERROR_NOT_EXIST_USER:              "用户不存在",
	ERROR_GET_USER_FAIL:               "获取用户失败",
	ERROR_GET_ALL_USER_FAIL:           "获取所有用户失败",
	ERROR_CLEAN_SOFT_DELETE_USER_FAIL: "清空所有软删除用户失败",
	ERROR_ADD_USER_FAIL:               "添加用户失败",
	ERROR_DELETE_USER_FAIL:            "删除用户失败",
	ERROR_CHECK_EXIST_USER_FAIL:       "验证已存在用户失败",
	ERROR_COUNT_USER_FAIL:             "统计用户数量失败",
	ERROR_CHECK_PASSWORD_FAIL:         "验证密码失败",
	ERROR_USER_NAME_OR_PASSWORD:       "用户名或者密码错误",
	ERROR_EXIST_USER_NAME:             "用户名已存在",
	ERROR_RESET_PASSWORD_FAIL:         "修改密码失败",
	ERROR_PASSWORD_SIZE:               "密码长度不能小于6位",

	ERROR_NOT_EXIST_API:        "不存在",
	ERROR_GET_API_FAIL:         "获取失败",
	ERROR_GET_ALL_API_FAIL:     "获取所有失败",
	ERROR_CHECK_EXIST_API_FAIL: "验证已存在失败",
	ERROR_COUNT_API_FAIL:       "统计数量失败",

	ERROR_REPEAT_LUCKY:        "正在抽奖，请稍后重试",
	ERROR_DAY_NOT_LUCKY_COUNT: "今日抽奖次数已用完",

	ERROR_CHECK_TOKEN_FAIL:  "Token鉴权失败",
	ERROR_TOKEN_TIMEOUT:     "Token已超时",
	ERROR_CREATE_TOKEN_FAIL: "Token生成失败",
	ERROR_TOKEN:             "Token错误",
	ERROR_PARSE_TOKEN_FAIL:  "Token解析失败",

	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",

	ERROR_GET_CAPTCHA_FAIL: "生成验证码失败",
	ERROR_CAPTCHA: "验证码错误",

	ERROR_GET_ROLE_FAIL: "获取角色信息失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
