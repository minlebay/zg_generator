package grpc_client

import (
	"fmt"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"zg_generator/pkg/message_v1"
)

type Client struct {
	Logger     *zap.Logger
	Config     *Config
	GrpcClient message.MessageRouterClient
	Conn       *grpc.ClientConn
}

func NewClient(logger *zap.Logger, config *Config) *Client {
	return &Client{
		Logger: logger,
		Config: config,
	}
}

func (r *Client) StartClient() {
	go func() {
		grpcTarget := fmt.Sprintf("%s", r.Config.RouterAddress)

		conn, err := grpc.NewClient(
			grpcTarget,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(grpcmiddleware.ChainUnaryClient(
				grpcprometheus.UnaryClientInterceptor,
			)),
			grpc.WithStreamInterceptor(grpcmiddleware.ChainStreamClient(
				grpcprometheus.StreamClientInterceptor,
			)),
		)
		if err != nil {
			r.Logger.Fatal(err.Error())
		}

		r.Conn = conn
		r.GrpcClient = message.NewMessageRouterClient(conn)

	}()
}

func (r *Client) StopClient() {
	r.Conn.Close()
}
