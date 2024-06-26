package grpc_client

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"zg_generator/pkg/message_v1/router"
)

type Client struct {
	Done       chan struct{}
	Logger     *zap.Logger
	Config     *Config
	GrpcClient router.MessageRouterClient
	Conn       *grpc.ClientConn
}

func NewClient(logger *zap.Logger, config *Config) *Client {
	return &Client{
		Done:   make(chan struct{}),
		Logger: logger,
		Config: config,
	}
}

func (r *Client) StartClient() {
	grpcTarget := fmt.Sprintf("%s", r.Config.RouterAddress)

	conn, err := grpc.NewClient(grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		r.Logger.Fatal(err.Error())
	}

	r.Conn = conn
	r.GrpcClient = router.NewMessageRouterClient(conn)
	r.Logger.Info("Client started")
}

func (r *Client) StopClient() {
	r.Conn.Close()
	r.Logger.Info("Client stopped")
	r.Done <- struct{}{}
}
