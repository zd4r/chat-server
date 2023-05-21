package app

import (
	"context"
	"log"

	chatV1 "github.com/zd4r/chat-server/internal/api/chat_v1"
	"github.com/zd4r/chat-server/internal/config"
	chatService "github.com/zd4r/chat-server/internal/service/chat"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	chatService chatService.Service

	chatImpl *chatV1.Implementation
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetGRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) GetChatService(_ context.Context) chatService.Service {
	if s.chatService == nil {
		s.chatService = chatService.NewService()
	}

	return s.chatService
}

func (s *serviceProvider) GetChatImpl(ctx context.Context) *chatV1.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chatV1.NewImplementation(s.GetChatService(ctx))
	}

	return s.chatImpl
}
