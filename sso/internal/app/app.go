package app

import (
	grpcapp "sso/sso/internal/app/grpc"
	"sso/sso/internal/logger"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *logger.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	//TODO ::

	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
