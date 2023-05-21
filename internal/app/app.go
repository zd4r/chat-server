package app

import (
	"context"
	"net"

	"github.com/zd4r/chat-server/internal/config"
	desc "github.com/zd4r/chat-server/pkg/chat_v1"
	"github.com/zd4r/chat-server/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type app struct {
	serviceProvider *serviceProvider

	grpcServer *grpc.Server
}

func NewApp(ctx context.Context) (*app, error) {
	var a app

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *app) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	err := a.runGRPCServer()
	if err != nil {
		return err
	}

	return nil
}

func (a *app) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *app) initServiceProvider(_ context.Context) error {
	a.serviceProvider = NewServiceProvider()

	return nil
}

func (a *app) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer()
	reflection.Register(a.grpcServer)

	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.GetChatImpl(ctx))

	return nil
}

func (a *app) runGRPCServer() error {
	lis, err := net.Listen("tcp", a.serviceProvider.GetGRPCConfig().Host())
	if err != nil {
		return err
	}

	return a.grpcServer.Serve(lis)
}
