package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	UNAUTHORIZED   = 401

	ERROR_EDIT_USER_FAIL              = 10101
	ERROR_NOT_EXIST_USER              = 10102
	ERROR_GET_USER_FAIL               = 10103
	ERROR_GET_ALL_USER_FAIL           = 10104
	ERROR_CLEAN_SOFT_DELETE_USER_FAIL = 10105
	ERROR_ADD_USER_FAIL               = 10106
	ERROR_DELETE_USER_FAIL            = 10107
	ERROR_CHECK_EXIST_USER_FAIL       = 10108
	ERROR_COUNT_USER_FAIL             = 10109
	ERROR_CHECK_PASSWORD_FAIL         = 10110
	ERROR_USER_NAME_OR_PASSWORD       = 10111
	ERROR_EXIST_USER_NAME             = 10112
	ERROR_RESET_PASSWORD_FAIL         = 10113
	ERROR_PASSWORD_SIZE               = 10114

	ERROR_NOT_EXIST_API        = 10202
	ERROR_GET_API_FAIL         = 10203
	ERROR_GET_ALL_API_FAIL     = 10204
	ERROR_CHECK_EXIST_API_FAIL = 10208
	ERROR_COUNT_API_FAIL       = 10209

	ERROR_REPEAT_LUCKY        = 10310
	ERROR_DAY_NOT_LUCKY_COUNT = 10311

	ERROR_CHECK_TOKEN_FAIL  = 10401
	ERROR_TOKEN_TIMEOUT     = 10402
	ERROR_TOKEN             = 10403
	ERROR_CREATE_TOKEN_FAIL = 10404
	ERROR_PARSE_TOKEN_FAIL  = 10405

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 10501
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 10502
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 10503

	ERROR_GET_CAPTCHA_FAIL = 10601
	ERROR_CAPTCHA = 10602

	ERROR_GET_ROLE_FAIL = 10701
)
