package component

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"opencsg.com/csghub-server/builder/memory"
	"opencsg.com/csghub-server/common/types"
)

type mockMemoryAdapter struct {
	mock.Mock
}

func (m *mockMemoryAdapter) GetModels(ctx context.Context) (*types.GetMemoryModelsResponse, error) {
	ret := m.Called(ctx)
	var r0 *types.GetMemoryModelsResponse
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*types.GetMemoryModelsResponse)
	}
	return r0, ret.Error(1)
}

func (m *mockMemoryAdapter) GetChatModel(ctx context.Context) (*types.MemoryChatModelConfig, error) {
	ret := m.Called(ctx)
	var r0 *types.MemoryChatModelConfig
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*types.MemoryChatModelConfig)
	}
	return r0, ret.Error(1)
}

func (m *mockMemoryAdapter) GetEmbeddingModel(ctx context.Context) (*types.MemoryEmbeddingModelConfig, error) {
	ret := m.Called(ctx)
	var r0 *types.MemoryEmbeddingModelConfig
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*types.MemoryEmbeddingModelConfig)
	}
	return r0, ret.Error(1)
}

func (m *mockMemoryAdapter) SetModels(ctx context.Context, req *types.SetMemoryModelsRequest) error {
	ret := m.Called(ctx, req)
	return ret.Error(0)
}

func (m *mockMemoryAdapter) SetChatModel(ctx context.Context, req *types.MemoryChatModelConfig) error {
	ret := m.Called(ctx, req)
	return ret.Error(0)
}

func (m *mockMemoryAdapter) SetEmbeddingModel(ctx context.Context, req *types.MemoryEmbeddingModelConfig) error {
	ret := m.Called(ctx, req)
	return ret.Error(0)
}

func (m *mockMemoryAdapter) DeleteChatModel(ctx context.Context) error {
	ret := m.Called(ctx)
	return ret.Error(0)
}

func (m *mockMemoryAdapter) DeleteEmbeddingModel(ctx context.Context) error {
	ret := m.Called(ctx)
	return ret.Error(0)
}

func (m *mockMemoryAdapter) CreateProject(ctx context.Context, req *types.CreateMemoryProjectRequest) (*types.MemoryProjectResponse, error) {
	ret := m.Called(ctx, req)
	var r0 *types.MemoryProjectResponse
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*types.MemoryProjectResponse)
	}
	return r0, ret.Error(1)
}

func (m *mockMemoryAdapter) GetProject(ctx context.Context, req *types.GetMemoryProjectRequest) (*types.MemoryProjectResponse, error) {
	ret := m.Called(ctx, req)
	var r0 *types.MemoryProjectResponse
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*types.MemoryProjectResponse)
	}
	return r0, ret.Error(1)
}

func (m *mockMemoryAdapter) ListProjects(ctx context.Context) ([]*types.MemoryProjectRef, error) {
	ret := m.Called(ctx)
	var r0 []*types.MemoryProjectRef
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*types.MemoryProjectRef)
	}
	return r0, ret.Error(1)
}

func (m *mockMemoryAdapter) DeleteProject(ctx context.Context, req *types.DeleteMemoryProjectRequest) error {
	ret := m.Called(ctx, req)
	return ret.Error(0)
}

func (m *mockMemoryAdapter) AddMemories(ctx context.Context, req *types.AddMemoriesRequest) (*types.AddMemoriesResponse, error) {
	ret := m.Called(ctx, req)
	var r0 *types.AddMemoriesResponse
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*types.AddMemoriesResponse)
	}
	return r0, ret.Error(1)
}

func (m *mockMemoryAdapter) SearchMemories(ctx context.Context, req *types.SearchMemoriesRequest) (*types.SearchMemoriesResponse, error) {
	ret := m.Called(ctx, req)
	var r0 *types.SearchMemoriesResponse
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*types.SearchMemoriesResponse)
	}
	return r0, ret.Error(1)
}

func (m *mockMemoryAdapter) ListMemories(ctx context.Context, req *types.ListMemoriesRequest) (*types.ListMemoriesResponse, error) {
	ret := m.Called(ctx, req)
	var r0 *types.ListMemoriesResponse
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*types.ListMemoriesResponse)
	}
	return r0, ret.Error(1)
}

func (m *mockMemoryAdapter) DeleteMemories(ctx context.Context, req *types.DeleteMemoriesRequest) error {
	ret := m.Called(ctx, req)
	return ret.Error(0)
}

func (m *mockMemoryAdapter) Health(ctx context.Context) (*types.MemoryHealthResponse, error) {
	ret := m.Called(ctx)
	var r0 *types.MemoryHealthResponse
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*types.MemoryHealthResponse)
	}
	return r0, ret.Error(1)
}

var _ memory.Adapter = (*mockMemoryAdapter)(nil)

func TestMemoryComponent_GetModels(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	resp := &types.GetMemoryModelsResponse{
		Chat:      &types.MemoryChatModelConfig{BaseURL: "http://chat", APIKey: "k", Model: "m"},
		Embedding: &types.MemoryEmbeddingModelConfig{BaseURL: "http://embed", APIKey: "k", Model: "e", Dimensions: 256},
	}
	client.On("GetModels", mock.Anything).Return(resp, nil)

	got, err := component.GetModels(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, resp, got)
}

func TestMemoryComponent_GetChatModel(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	resp := &types.MemoryChatModelConfig{BaseURL: "http://chat", APIKey: "k", Model: "m"}
	client.On("GetChatModel", mock.Anything).Return(resp, nil)

	got, err := component.GetChatModel(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, resp, got)
}

func TestMemoryComponent_GetEmbeddingModel(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	resp := &types.MemoryEmbeddingModelConfig{BaseURL: "http://embed", APIKey: "k", Model: "e", Dimensions: 256}
	client.On("GetEmbeddingModel", mock.Anything).Return(resp, nil)

	got, err := component.GetEmbeddingModel(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, resp, got)
}

func TestMemoryComponent_SetModels(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	req := &types.SetMemoryModelsRequest{
		Chat:      &types.MemoryChatModelConfig{BaseURL: "http://chat", APIKey: "k", Model: "m"},
		Embedding: &types.MemoryEmbeddingModelConfig{BaseURL: "http://embed", APIKey: "k", Model: "e", Dimensions: 256},
	}
	client.On("SetModels", mock.Anything, req).Return(nil)

	err := component.SetModels(context.Background(), req)
	assert.NoError(t, err)
}

func TestMemoryComponent_SetChatModel(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	req := &types.MemoryChatModelConfig{BaseURL: "http://chat", APIKey: "k", Model: "m"}
	client.On("SetChatModel", mock.Anything, req).Return(nil)

	err := component.SetChatModel(context.Background(), req)
	assert.NoError(t, err)
}

func TestMemoryComponent_SetEmbeddingModel(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	req := &types.MemoryEmbeddingModelConfig{BaseURL: "http://embed", APIKey: "k", Model: "e", Dimensions: 256}
	client.On("SetEmbeddingModel", mock.Anything, req).Return(nil)

	err := component.SetEmbeddingModel(context.Background(), req)
	assert.NoError(t, err)
}

func TestMemoryComponent_DeleteChatModel(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	client.On("DeleteChatModel", mock.Anything).Return(nil)

	err := component.DeleteChatModel(context.Background())
	assert.NoError(t, err)
}

func TestMemoryComponent_DeleteEmbeddingModel(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	client.On("DeleteEmbeddingModel", mock.Anything).Return(nil)

	err := component.DeleteEmbeddingModel(context.Background())
	assert.NoError(t, err)
}

func TestMemoryComponent_CreateProject(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	req := &types.CreateMemoryProjectRequest{OrgID: "org", ProjectID: "proj"}
	resp := &types.MemoryProjectResponse{OrgID: "org", ProjectID: "proj"}
	client.On("CreateProject", mock.Anything, req).Return(resp, nil)

	got, err := component.CreateProject(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp, got)
}

func TestMemoryComponent_Capabilities(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	caps := component.Capabilities()
	assert.True(t, caps.SupportsProject)
	assert.True(t, caps.SupportsList)
	assert.False(t, caps.SupportsMetrics)
}

func TestMemoryComponent_GetProject(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	req := &types.GetMemoryProjectRequest{OrgID: "org", ProjectID: "proj"}
	resp := &types.MemoryProjectResponse{OrgID: "org", ProjectID: "proj"}
	client.On("GetProject", mock.Anything, req).Return(resp, nil)

	got, err := component.GetProject(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp, got)
}

func TestMemoryComponent_ListProjects(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	resp := []*types.MemoryProjectRef{{OrgID: "org", ProjectID: "proj"}}
	client.On("ListProjects", mock.Anything).Return(resp, nil)

	got, err := component.ListProjects(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, resp, got)
}

func TestMemoryComponent_DeleteProject(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	req := &types.DeleteMemoryProjectRequest{OrgID: "org", ProjectID: "proj"}
	client.On("DeleteProject", mock.Anything, req).Return(nil)

	err := component.DeleteProject(context.Background(), req)
	assert.NoError(t, err)
}

func TestMemoryComponent_AddMemories(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	req := &types.AddMemoriesRequest{
		OrgID:     "org",
		ProjectID: "proj",
		Messages:  []types.MemoryMessage{{Content: "hello"}},
	}
	resp := &types.AddMemoriesResponse{Created: []types.MemoryMessage{{UID: "e_1", Content: "hello"}}}
	client.On("AddMemories", mock.Anything, req).Return(resp, nil)

	got, err := component.AddMemories(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp, got)
}

func TestMemoryComponent_SearchMemories(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	req := &types.SearchMemoriesRequest{ContentQuery: "query"}
	resp := &types.SearchMemoriesResponse{Status: 0, Content: []types.MemoryMessage{{UID: "e_1"}}}
	client.On("SearchMemories", mock.Anything, req).Return(resp, nil)

	got, err := component.SearchMemories(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp, got)
}

func TestMemoryComponent_ListMemories(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	req := &types.ListMemoriesRequest{OrgID: "org", ProjectID: "proj"}
	resp := &types.ListMemoriesResponse{Status: 0, Content: []types.MemoryMessage{{UID: "e_1"}}}
	client.On("ListMemories", mock.Anything, req).Return(resp, nil)

	got, err := component.ListMemories(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp, got)
}

func TestMemoryComponent_DeleteMemories(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	req := &types.DeleteMemoriesRequest{OrgID: "org", ProjectID: "proj", UID: "e_1"}
	client.On("DeleteMemories", mock.Anything, req).Return(nil)

	err := component.DeleteMemories(context.Background(), req)
	assert.NoError(t, err)
}

func TestMemoryComponent_Health(t *testing.T) {
	client := &mockMemoryAdapter{}
	component := newMemoryComponent(client)

	resp := &types.MemoryHealthResponse{Status: "healthy", Service: "memmachine"}
	client.On("Health", mock.Anything).Return(resp, nil)

	got, err := component.Health(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, resp, got)
}
