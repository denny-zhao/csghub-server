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
