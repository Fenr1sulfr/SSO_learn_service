package grpcapp

import (
	"fmt"
	"net"
	authgrpc "sso/sso/internal/grpc/auth"
	"sso/sso/internal/logger"
	"strconv"

	"google.golang.org/grpc"
)

type App struct {
	log        *logger.Logger
	gRPCServer *grpc.Server
	port       int
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}
func (a *App) Run() error {
	const op = "grpcapp.run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}
	infoMap := make(map[string]string)
	infoMap["place"] = op
	infoMap["port"] = l.Addr().String()
	a.log.PrintInfo("gRPC server started", infoMap)
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s,%w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.stop"
	infoMap := make(map[string]string)
	infoMap["place"] = op
	infoMap["port"] = strconv.Itoa(a.port)

	a.log.PrintInfo("stopping gRPC server", infoMap)
	a.gRPCServer.GracefulStop()
}

func New(log *logger.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}
