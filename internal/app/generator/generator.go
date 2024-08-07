package generator

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sync"
	"time"
	"zg_generator/internal/app/grpc_client"
	"zg_generator/internal/app/telemetry"
	"zg_generator/pkg/message_v1"
)

type Generator struct {
	Done    chan struct{}
	Logger  *zap.Logger
	Config  *Config
	Client  *grpc_client.Client
	Metrics *telemetry.Metrics
	wg      sync.WaitGroup
}

func NewGenerator(
	logger *zap.Logger,
	config *Config,
	client *grpc_client.Client,
	metrics *telemetry.Metrics,
) *Generator {
	return &Generator{
		Done:    make(chan struct{}),
		Logger:  logger,
		Config:  config,
		Client:  client,
		Metrics: metrics,
	}
}

func (g *Generator) StartGenerator(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(time.Duration(g.Config.Interval) * time.Second)
		for {
			select {
			case <-ticker.C:
				for i := 0; i < g.Config.Count; i++ {
					go g.GenerateMessage(context.Background())
				}
			case <-g.Done:
				ticker.Stop()
				return
			default:
				continue
			}
		}
	}()
}

func (g *Generator) StopGenerator(ctx context.Context) {
	g.wg.Wait()
	g.Done <- struct{}{}
}

func (g *Generator) GenerateMessage(ctx context.Context) {
	g.wg.Add(1)
	defer g.wg.Done()

	in := &message.Message{
		Uuid:        uuid.NewString(),
		ContentType: "text/plain",
		MessageContent: &message.MessageContent{
			SendAt:   timestamppb.New(time.Now()),
			Provider: "default provider",
			Consumer: "default consumer",
			Title:    "default title",
			Content:  "this is default content",
		},
	}

	resp, err := g.Client.GrpcClient.ReceiveMessage(ctx, in)
	if err != nil {
		g.Logger.Error("failed to send message", zap.Error(err))
		return
	}

	g.Metrics.IncrementRequestCounter()
	g.Logger.Info("message sent", zap.Bool("success", resp.Success))
}
