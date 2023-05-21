package chat_v1

import (
	"github.com/zd4r/chat-server/internal/service/chat"
	desc "github.com/zd4r/chat-server/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server

	chatService chat.Service
}

func NewImplementation(chatService chat.Service) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
