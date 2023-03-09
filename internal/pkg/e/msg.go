package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	TOKEN_INVALID:  "Token已失效，请重新登录",

	API_SIGN_FAILED: "API 签名验证失败",
	EMPTY:           "没有记录",
	NOT_LOGIN:       "没有登录",
	NOT_PERMISSION:  "没有权限",
	CUSTOMER_EMPTY:  "没有该客户记录",

	ERROR_EXISTS:        "数据已存在",
	ERROR_GET_FAILED:    "数据获取失败",
	ERROR_CREATE_FAILED: "数据添加失败",
	ERROR_DELETE_FAILED: "数据删除失败",
	ERROR_UPDATE_FAILED: "数据更新失败",
	ERROR_UPLOAD_FAILED: "上传失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
