package generator

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"zg_generator/internal/app/grpc_client"
	"zg_generator/internal/model"
	"zg_generator/pkg/message_v1/router"
)

type Generator struct {
	Done   chan struct{}
	Logger *zap.Logger
	Config *Config
	Client *grpc_client.Client
}

func NewGenerator(logger *zap.Logger, config *Config, client *grpc_client.Client) *Generator {
	return &Generator{
		Done:   make(chan struct{}),
		Logger: logger,
		Config: config,
		Client: client,
	}
}

func (g *Generator) StartGenerator() {
	g.Logger.Info("Generator started")

	ticker := time.NewTicker(time.Duration(g.Config.Interval) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				for i := 0; i < g.Config.Count; i++ {
					go g.GenerateMessage()
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

func (g *Generator) StopGenerator() {
	g.Logger.Info("Generator stopped")
	g.Done <- struct{}{}
}

func (g *Generator) GenerateMessage() {
	// TODO implement the business logic of GenerateMessage
	m := &model.Message{
		UUID:        uuid.NewString(),
		ContentType: "text/plain",
		MessageContent: model.MessageContent{
			SendAt:   time.Now(),
			Provider: "default provider",
			Consumer: "default consumer",
			Title:    "default title",
			Content:  "this is default content",
		},
	}

	in := &router.Message{
		Uuid:        m.UUID,
		ContentType: m.ContentType,
		MessageContent: &router.MessageContent{
			SendAt:   timestamppb.New(m.MessageContent.SendAt),
			Provider: m.MessageContent.Provider,
			Consumer: m.MessageContent.Consumer,
			Title:    m.MessageContent.Title,
			Content:  m.MessageContent.Content,
		},
	}

	resp, err := g.Client.GrpcClient.ReceiveMessage(context.Background(), in)
	if err != nil {
		g.Logger.Error("failed to send message", zap.Error(err))
		return
	}

	g.Logger.Info("message sent", zap.Bool("success", resp.Success))
}
