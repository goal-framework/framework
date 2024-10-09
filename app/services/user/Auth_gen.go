// Code generated by goal-cli. DO NOT EDIT.
// versions:
// 	goal-cli v0.5.24
// 	go       go1.23.2
//
// updated_at: 2024-10-09 19:07:27
// source: pro/user.proto
// 
package user

import (
	user "github.com/goal-web/goal/app/requests/user"
	user0 "github.com/goal-web/goal/app/results/user"
)

var AuthServiceDefine AuthServiceStatic

type AuthServiceStatic struct {
	LoginByWxAppCode func(req *user.LoginByWxCodeReq) (*user0.LoginByWxResult, error)
	LoginByWxAppInfo func(req *user.LoginByWxInfoReq) (*user0.LoginByWxResult, error)
}

func AuthServiceLoginByWxAppCode(req *user.LoginByWxCodeReq) (*user0.LoginByWxResult, error) {
	if AuthServiceDefine.LoginByWxAppCode != nil {
		return AuthServiceDefine.LoginByWxAppCode(req)
	}
	return nil, nil
}

func AuthServiceLoginByWxAppInfo(req *user.LoginByWxInfoReq) (*user0.LoginByWxResult, error) {
	if AuthServiceDefine.LoginByWxAppInfo != nil {
		return AuthServiceDefine.LoginByWxAppInfo(req)
	}
	return nil, nil
}
