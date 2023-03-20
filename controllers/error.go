package controllers

const (
	AUTH             = "auth"
	DEFAULT_TIPS     = "该项不能为空!"
	DEFAULT_MIN_SIZE = 1
	//the base
	RES_OK              = 0
	FALSE               = -1
	TOKEN_PRODUCE_FALSE = -2
	TOKEN_INVALIDATE    = -3
	USER_IS_NULL 	= -4
	VALIDATE_FAIL = -5

	//the systen:10~300
	MAC_ISNULL     = -10
	DEVICE_ISNULL  = -11
	GETINT_FALSE   = -12
	GETINT64_FALSE = -13

	//the user:301~500
	DEFAULT_ISNULL    = -301
	DEFAULT_SIZE_LOW  = -302
	PASSWORD_DIFFER   = -303	
	USER_NOT_ACTIVE   = -304
	USER_FORBIDDEN    = -305
	USER_FORBIDDEN_WORDS=-306
	USER_EXCEPTION=-307
	USER_PASSWORD_ERR = -308
	USER_ALREADY_EXIISTA=-309
	USER_SET_INFO=-310
	USER_NOT_EXIST=-311//用户不存在

	//the db:501~800
	DB_INSERT_FALSE = 501
	DB_DELETE_FALSE = 502
	DB_UPDATE_FALSE = 503
	DB_QUERY_FALSE  = 503
	//the net:801~1200
	//the other:1201~...

	REG=1
	BIND=11
	LOGIN=2
	CHANNEL=3
	BLOG=4

	DEFAULT_INTRO="记忆美好生活，从梵记开始..."
)

var msg = map[int]interface{}{
	FALSE:"数据为空!",
	//the system:10~300
	MAC_ISNULL:    "Mac地址不能为空!",
	DEVICE_ISNULL: "设备名称不能为空!",
	//the user:301~500
	PASSWORD_DIFFER:   "密码不一致!",
	USER_NOT_ACTIVE:   "该账号未激活!",
	USER_FORBIDDEN:    "该账号已被禁用!",
	USER_FORBIDDEN_WORDS:    "该账号已被禁言!",
	USER_EXCEPTION:    "该账号异常，请联系客户!",
	USER_PASSWORD_ERR: "输入密码错误!",
	USER_IS_NULL:"用户不存在!",
	VALIDATE_FAIL:"验证码错误",
	USER_ALREADY_EXIISTA:"用户已存在!",
	USER_SET_INFO:"用户还未设置信息",
	USER_NOT_EXIST:"用户不存在!",
	//the db:501~800

	//the net:801~1200

	//the other:1201~...
}

type ErrorController struct {
	BaseController
}

func (this *ErrorController) Error404() {
	this.Data["content"] = "正在开发中..."
	this.false(-1, "404错误")
}

func (this *ErrorController) Error501() {
	this.Data["content"] = "server error"
	this.false(-1, "501错误")
}

func (this *ErrorController) ErrorDb() {
	this.Data["content"] = "database is now down"
	this.false(-1, "db操作错误")
}
