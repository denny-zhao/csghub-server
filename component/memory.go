package component

import (
	"context"
	"fmt"
	"time"

	"opencsg.com/csghub-server/builder/memory"
	_ "opencsg.com/csghub-server/builder/memory/memmachine"
	"opencsg.com/csghub-server/builder/rpc"
	"opencsg.com/csghub-server/common/config"
	"opencsg.com/csghub-server/common/types"
)

type MemoryComponent interface {
	Capabilities() types.MemoryCapabilities
	GetModels(ctx context.Context) (*types.GetMemoryModelsResponse, error)
	GetChatModel(ctx context.Context) (*types.MemoryChatModelConfig, error)
	GetEmbeddingModel(ctx context.Context) (*types.MemoryEmbeddingModelConfig, error)
	SetModels(ctx context.Context, req *types.SetMemoryModelsRequest) error
	SetChatModel(ctx context.Context, req *types.MemoryChatModelConfig) error
	SetEmbeddingModel(ctx context.Context, req *types.MemoryEmbeddingModelConfig) error
	DeleteChatModel(ctx context.Context) error
	DeleteEmbeddingModel(ctx context.Context) error
	CreateProject(ctx context.Context, req *types.CreateMemoryProjectRequest) (*types.MemoryProjectResponse, error)
	GetProject(ctx context.Context, req *types.GetMemoryProjectRequest) (*types.MemoryProjectResponse, error)
	ListProjects(ctx context.Context) ([]*types.MemoryProjectRef, error)
	DeleteProject(ctx context.Context, req *types.DeleteMemoryProjectRequest) error
	AddMemories(ctx context.Context, req *types.AddMemoriesRequest) (*types.AddMemoriesResponse, error)
	SearchMemories(ctx context.Context, req *types.SearchMemoriesRequest) (*types.SearchMemoriesResponse, error)
	ListMemories(ctx context.Context, req *types.ListMemoriesRequest) (*types.ListMemoriesResponse, error)
	DeleteMemories(ctx context.Context, req *types.DeleteMemoriesRequest) error
	Health(ctx context.Context) (*types.MemoryHealthResponse, error)
}

type memoryComponentImpl struct {
	client       memory.Adapter
	capabilities types.MemoryCapabilities
}

func NewMemoryComponent(cfg *config.Config) (MemoryComponent, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}
	memCfg := cfg.Memory
	endpoint := fmt.Sprintf("%s:%d", memCfg.Host, memCfg.Port)
	adapter, err := memory.NewAdapter(memCfg.Backend, endpoint, memCfg.BasePath, buildMemoryOpts(memCfg)...)
	if err != nil {
		return nil, err
	}
	if memCfg.TimeoutSeconds > 0 {
		if setter, ok := adapter.(memory.TimeoutSetter); ok {
			setter.WithTimeout(time.Duration(memCfg.TimeoutSeconds) * time.Second)
		}
	}
	if memCfg.RetryCount > 0 {
		if setter, ok := adapter.(memory.RetrySetter); ok {
			setter.WithRetry(uint(memCfg.RetryCount))
		}
	}
	if memCfg.RetryDelayMillis > 0 {
		if setter, ok := adapter.(memory.DelaySetter); ok {
			setter.WithDelay(time.Duration(memCfg.RetryDelayMillis) * time.Millisecond)
		}
	}
	return newMemoryComponent(adapter), nil
}

func buildMemoryOpts(cfg config.MemoryConfig) []rpc.RequestOption {
	var opts []rpc.RequestOption
	if cfg.ApiKey != "" {
		opts = append(opts, rpc.AuthWithApiKey(cfg.ApiKey))
	}
	return opts
}

func newMemoryComponent(client memory.Adapter) *memoryComponentImpl {
	return &memoryComponentImpl{
		client: client,
		capabilities: types.MemoryCapabilities{
			SupportsProject:     true,
			SupportsList:        true,
			SupportsMetrics:     false,
			SupportsHealthCheck: true,
		},
	}
}

func (c *memoryComponentImpl) Capabilities() types.MemoryCapabilities {
	return c.capabilities
}

func (c *memoryComponentImpl) GetModels(ctx context.Context) (*types.GetMemoryModelsResponse, error) {
	return c.client.GetModels(ctx)
}

func (c *memoryComponentImpl) GetChatModel(ctx context.Context) (*types.MemoryChatModelConfig, error) {
	return c.client.GetChatModel(ctx)
}

func (c *memoryComponentImpl) GetEmbeddingModel(ctx context.Context) (*types.MemoryEmbeddingModelConfig, error) {
	return c.client.GetEmbeddingModel(ctx)
}

func (c *memoryComponentImpl) SetModels(ctx context.Context, req *types.SetMemoryModelsRequest) error {
	return c.client.SetModels(ctx, req)
}

func (c *memoryComponentImpl) SetChatModel(ctx context.Context, req *types.MemoryChatModelConfig) error {
	return c.client.SetChatModel(ctx, req)
}

func (c *memoryComponentImpl) SetEmbeddingModel(ctx context.Context, req *types.MemoryEmbeddingModelConfig) error {
	return c.client.SetEmbeddingModel(ctx, req)
}

func (c *memoryComponentImpl) DeleteChatModel(ctx context.Context) error {
	return c.client.DeleteChatModel(ctx)
}

func (c *memoryComponentImpl) DeleteEmbeddingModel(ctx context.Context) error {
	return c.client.DeleteEmbeddingModel(ctx)
}

func (c *memoryComponentImpl) CreateProject(ctx context.Context, req *types.CreateMemoryProjectRequest) (*types.MemoryProjectResponse, error) {
	return c.client.CreateProject(ctx, req)
}

func (c *memoryComponentImpl) GetProject(ctx context.Context, req *types.GetMemoryProjectRequest) (*types.MemoryProjectResponse, error) {
	return c.client.GetProject(ctx, req)
}

func (c *memoryComponentImpl) ListProjects(ctx context.Context) ([]*types.MemoryProjectRef, error) {
	return c.client.ListProjects(ctx)
}

func (c *memoryComponentImpl) DeleteProject(ctx context.Context, req *types.DeleteMemoryProjectRequest) error {
	return c.client.DeleteProject(ctx, req)
}

func (c *memoryComponentImpl) AddMemories(ctx context.Context, req *types.AddMemoriesRequest) (*types.AddMemoriesResponse, error) {
	return c.client.AddMemories(ctx, req)
}

func (c *memoryComponentImpl) SearchMemories(ctx context.Context, req *types.SearchMemoriesRequest) (*types.SearchMemoriesResponse, error) {
	return c.client.SearchMemories(ctx, req)
}

func (c *memoryComponentImpl) ListMemories(ctx context.Context, req *types.ListMemoriesRequest) (*types.ListMemoriesResponse, error) {
	return c.client.ListMemories(ctx, req)
}

func (c *memoryComponentImpl) DeleteMemories(ctx context.Context, req *types.DeleteMemoriesRequest) error {
	return c.client.DeleteMemories(ctx, req)
}

func (c *memoryComponentImpl) Health(ctx context.Context) (*types.MemoryHealthResponse, error) {
	return c.client.Health(ctx)
}
