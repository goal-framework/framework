// Code generated by goal-cli. DO NOT EDIT.
// versions:
// 	goal-cli v0.5.24
// 	go       go1.23.2
//
// updated_at: 2024-10-09 19:07:27
// source: pro/common.proto
// 
package enums

type Code int

const (

	// CodeSuccess
	// @msg:成功
	CodeSuccess Code = 0

	// CodeParseReqErr
	// @msg:参数解析失败
	CodeParseReqErr Code = 10400

	// CodeUnauthorized
	// @msg:未登录
	CodeUnauthorized Code = 10401

	// CodeForbidden
	// @msg:没有权限
	CodeForbidden Code = 100403

	// CodeNotFound
	// @msg:找不到
	CodeNotFound Code = 100404

	// CodeBizErr
	// @msg:业务错误
	CodeBizErr  Code = 10201
	CodeUnknown Code = -1000
)

func (item Code) String() string {
	switch item {
	case CodeSuccess:
		return "Success"
	case CodeParseReqErr:
		return "ParseReqErr"
	case CodeUnauthorized:
		return "Unauthorized"
	case CodeForbidden:
		return "Forbidden"
	case CodeNotFound:
		return "NotFound"
	case CodeBizErr:
		return "BizErr"
	default:
		return "Unknown"
	}
}

func (item Code) Message() string {
	switch item {
	case CodeSuccess:
		return "成功"
	case CodeParseReqErr:
		return "参数解析失败"
	case CodeUnauthorized:
		return "未登录"
	case CodeForbidden:
		return "没有权限"
	case CodeNotFound:
		return "找不到"
	case CodeBizErr:
		return "业务错误"
	default:
		return "Unknown"
	}
}

func ParseCodeFromString(msg string) Code {
	switch msg {
	case "Success":
		return CodeSuccess
	case "ParseReqErr":
		return CodeParseReqErr
	case "Unauthorized":
		return CodeUnauthorized
	case "Forbidden":
		return CodeForbidden
	case "NotFound":
		return CodeNotFound
	case "BizErr":
		return CodeBizErr
	default:
		return CodeUnknown
	}
}
