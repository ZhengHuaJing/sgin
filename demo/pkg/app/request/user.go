package request

type LoginUserForm struct {
	UserName  string `json:"user_name" form:"user_name" valid:"Match(/^[a-zA-Z0-9_-]{4,16}$/)"` // 用户名
	Password  string `json:"password" form:"password" valid:"MinSize(6)"`                       // 密码
	Captcha   string `json:"captcha" form:"captcha" valid:"Required"`                           // 验证码
	CaptchaID string `json:"captcha_id" form:"captcha_id" valid:"Required"`                     // 验证码ID
}

type AddUserForm struct {
	UserName string `json:"user_name" form:"user_name" valid:"Match(/^[a-zA-Z0-9_-]{4,16}$/)"` // 用户名
	Password string `json:"password" form:"password" valid:"MinSize(6)"`                       // 密码
}

type UpdateUserForm struct {
	ID       int    `json:"-" form:"-" valid:"Min(1)"`                     // ID
	Password string `json:"password" form:"password"`                      // 密码
	RealName string `json:"real_name" form:"real_name" valid:"MinSize(2)"` // 联系人
	Mobile   string `json:"mobile" form:"mobile" valid:"Mobile"`           // 手机号
	Address  string `json:"address" form:"address" valid:"MinSize(6)"`     // 地址
	Gender   string `json:"gender" form:"Address" valid:"Length(1)"`       // 性别
}
