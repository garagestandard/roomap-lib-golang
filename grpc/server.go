package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

func SetupLogger() (*zap.Logger, grpc_zap.Option) {
	zap, _ := zap.NewProduction()
	zap_opt := grpc_zap.WithLevels(
		func(c codes.Code) zapcore.Level {
			var l zapcore.Level
			switch c {
			case codes.OK:
				l = zapcore.InfoLevel

			case codes.Internal:
				l = zapcore.ErrorLevel

			default:
				l = zapcore.DebugLevel
			}
			return l
		},
	)
	return zap, zap_opt
}

func CreateUnaryServer() *grpc.Server {
	zap, zap_opt := SetupLogger()

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap, zap_opt),
		),
	)
	return grpcServer
}

func StartUnaryServer(grpcServer *grpc.Server) {

	port := os.Getenv("GRPC_LISTEN_PORT")
	listenPort, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		reflection.Register(grpcServer)
		grpcServer.Serve(listenPort)
	}()
	log.Printf("start gRPC server port: %v", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC s...")
	grpcServer.GracefulStop()
}
