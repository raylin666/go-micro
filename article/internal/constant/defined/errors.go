package defined

import api "mt/api/v1"

var (
	/* 系统相关 */
	ErrorUnknownError           = api.ErrorUnknownError("未知错误")
	ErrorServerError            = api.ErrorServerError("服务异常")
	ErrorDataValidateError      = api.ErrorDataValidateError("数据校验失败")
	ErrorDataSelectError        = api.ErrorDataSelectError("数据查询失败")
	ErrorDataAlreadyExists      = api.ErrorDataAlreadyExists("数据已存在")
	ErrorDataNotFound           = api.ErrorDataNotFound("数据不存在")
	ErrorDataAddError           = api.ErrorDataAddError("新增数据失败")
	ErrorDataUpdateError        = api.ErrorDataUpdateError("更新数据失败")
	ErrorDataDeleteError        = api.ErrorDataDeleteError("数据删除失败")
	ErrorDataResourceNotFound   = api.ErrorDataResourceNotFound("数据资源不存在")
	ErrorDataUpdateFieldError   = api.ErrorDataUpdateFieldError("数据属性更新失败")
	ErrorIdInvalidValueError    = api.ErrorIdInvalidValueError("无效ID值")
	ErrorCommandInvalidNotFound = api.ErrorCommandInvalidNotFound("无效的执行指令")
	ErrorNotLoginError          = api.ErrorNotLoginError("请先登录后再操作")
	ErrorNotVisitAuth		    = api.ErrorNotVisitAuth("没有访问权限")
)
