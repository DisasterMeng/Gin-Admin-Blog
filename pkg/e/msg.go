package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_NOT_EXIST_CATEGORY:      "类别不存在",
	ERROR_NOT_EXIST_CATEGORY_FAIL: "获取类别失败",
	ERROR_DELETE_CATEGORY_FAIL:    "删除类别失败",
	ERROR_ADD_CATEGORY_FAIL:       "添加类别失败",
	ERROR_UPDATE_CATEGORY_FAIL:    "更新类别失败",

	ERROR_NOT_EXIST_TAG:      "标签不存在",
	ERROR_NOT_EXIST_TAG_FAIL: "获取标签失败",
	ERROR_DELETE_TAG_FAIL:    "删除标签失败",
	ERROR_ADD_TAG_FAIL:       "添加标签失败",
	ERROR_UPDATE_TAG_FAIL:    "更新标签失败",

	ERROR_NOT_EXIST_FRIEND:   "友链不存在",
	ERROR_ADD_FRIEND_FAIL:    "添加友链失败",
	ERROR_DELETE_FRIEND_FAIL: "删除友链失败",

	ERROR_ADD_BLOG_FAIL:    "添加Blog失败",
	ERROR_DELETE_BLOG_FAIL: "删除Blog失败",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
