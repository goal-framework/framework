// Code generated by goal-cli. DO NOT EDIT.
// versions:
// 	goal-cli v0.5.24
// 	go       go1.23.2
//
// updated_at: 2024-10-09 02:33:34
// source: pro/user.proto
// 
package user

import (
	"github.com/goal-web/contracts"
	models "github.com/goal-web/goal/app/models"
)

type LoginByWxResult struct {
	Token      string            `json:"token"`
	User       *models.UserModel `json:"user"`
	SessionKey string            `json:"sessionKey"`
}

func (result *LoginByWxResult) ToFields() contracts.Fields {
	fields := contracts.Fields{
		"token":      result.Token,
		"user":       result.User.ToFields(),
		"sessionKey": result.SessionKey,
	}

	return fields
}
