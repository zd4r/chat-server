package auth

import (
	"context"

	UserV1 "github.com/zd4r/auth/pkg/user_v1"
	converter "github.com/zd4r/chat-server/internal/converter/user"
	model "github.com/zd4r/chat-server/internal/model/user"
)

var _ Client = (*client)(nil)

type Client interface {
	Create(ctx context.Context, user *model.User) error
}

type client struct {
	userClient UserV1.UserV1Client
}

func NewClient(c UserV1.UserV1Client) *client {
	return &client{
		userClient: c,
	}
}

func (c *client) Create(ctx context.Context, user *model.User) error {
	_, err := c.userClient.Create(ctx, &UserV1.CreateRequest{
		User:            converter.ToUserDesc(user),
		PasswordConfirm: user.PasswordConfirm,
	})
	if err != nil {
		return err
	}

	return nil
}
