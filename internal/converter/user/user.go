package user

import (
	userV1 "github.com/zd4r/auth/pkg/user_v1"
	model "github.com/zd4r/chat-server/internal/model/user"
)

func ToUserDesc(user *model.User) *userV1.User {
	return &userV1.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     userV1.RoleInfo(userV1.RoleInfo_value[user.Role]),
	}
}
