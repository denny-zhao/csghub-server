package handler

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"opencsg.com/csghub-server/api/httpbase"
	"opencsg.com/csghub-server/common/config"
	"opencsg.com/csghub-server/common/errorx"
	"opencsg.com/csghub-server/common/types"
	"opencsg.com/csghub-server/component"
)

type MemoryHandler struct {
	memory component.MemoryComponent
}

func NewMemoryHandler(cfg *config.Config) (*MemoryHandler, error) {
	memoryComp, err := component.NewMemoryComponent(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create memory component: %w", err)
	}
	return &MemoryHandler{
		memory: memoryComp,
	}, nil
}

func (h *MemoryHandler) CreateProject(ctx *gin.Context) {
	var req types.CreateMemoryProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "bad request format", "error", err)
		httpbase.BadRequestWithExt(ctx, errorx.ReqBodyFormat(err, nil))
		return
	}
	resp, err := h.memory.CreateProject(ctx.Request.Context(), &req)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to create memory project", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, filterProjectResponse(resp))
}

func (h *MemoryHandler) GetProject(ctx *gin.Context) {
	var req types.GetMemoryProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "bad request format", "error", err)
		httpbase.BadRequestWithExt(ctx, errorx.ReqBodyFormat(err, nil))
		return
	}
	resp, err := h.memory.GetProject(ctx.Request.Context(), &req)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to get memory project", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, filterProjectResponse(resp))
}

func (h *MemoryHandler) ListProjects(ctx *gin.Context) {
	resp, err := h.memory.ListProjects(ctx.Request.Context())
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list memory projects", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, filterProjectListResponse(resp))
}

func (h *MemoryHandler) DeleteProject(ctx *gin.Context) {
	var req types.DeleteMemoryProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "bad request format", "error", err)
		httpbase.BadRequestWithExt(ctx, errorx.ReqBodyFormat(err, nil))
		return
	}
	if req.OrgID == "" || req.ProjectID == "" {
		httpbase.BadRequestWithExt(ctx, errorx.ReqParamInvalid(
			fmt.Errorf("org_id and project_id are required"),
			errorx.Ctx().Set("field", "org_id,project_id"),
		))
		return
	}
	if err := h.memory.DeleteProject(ctx.Request.Context(), &req); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to delete memory project", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, gin.H{"deleted": true})
}

func (h *MemoryHandler) AddMemories(ctx *gin.Context) {
	var req types.AddMemoriesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "bad request format", "error", err)
		httpbase.BadRequestWithExt(ctx, errorx.ReqBodyFormat(err, nil))
		return
	}
	for i, msg := range req.Messages {
		if msg.Scopes != nil && (msg.Scopes.AgentID != "" || msg.Scopes.OrgID != "" || msg.Scopes.ProjectID != "" || msg.Scopes.SessionID != "") {
			httpbase.BadRequestWithExt(ctx, errorx.ReqParamInvalid(
				fmt.Errorf("message scopes must not be set; use request-level scope fields"),
				errorx.Ctx().Set("field", fmt.Sprintf("messages[%d].scopes", i)),
			))
			return
		}
	}
	resp, err := h.memory.AddMemories(ctx.Request.Context(), &req)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to add memories", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, resp)
}

func filterProjectResponse(resp *types.MemoryProjectResponse) *types.MemoryProjectResponse {
	if resp == nil {
		return nil
	}
	return &types.MemoryProjectResponse{
		OrgID:       resp.OrgID,
		ProjectID:   resp.ProjectID,
		Description: resp.Description,
	}
}

func filterProjectListResponse(items []*types.MemoryProjectRef) []*types.MemoryProjectRef {
	if len(items) == 0 {
		return items
	}
	out := make([]*types.MemoryProjectRef, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}
		out = append(out, &types.MemoryProjectRef{
			OrgID:     item.OrgID,
			ProjectID: item.ProjectID,
		})
	}
	return out
}

func (h *MemoryHandler) SearchMemories(ctx *gin.Context) {
	var req types.SearchMemoriesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "bad request format", "error", err)
		httpbase.BadRequestWithExt(ctx, errorx.ReqBodyFormat(err, nil))
		return
	}
	resp, err := h.memory.SearchMemories(ctx.Request.Context(), &req)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to search memories", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, resp)
}

func (h *MemoryHandler) ListMemories(ctx *gin.Context) {
	var req types.ListMemoriesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "bad request format", "error", err)
		httpbase.BadRequestWithExt(ctx, errorx.ReqBodyFormat(err, nil))
		return
	}
	resp, err := h.memory.ListMemories(ctx.Request.Context(), &req)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list memories", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, resp)
}

func (h *MemoryHandler) DeleteMemories(ctx *gin.Context) {
	var req types.DeleteMemoriesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "bad request format", "error", err)
		httpbase.BadRequestWithExt(ctx, errorx.ReqBodyFormat(err, nil))
		return
	}
	if err := h.memory.DeleteMemories(ctx.Request.Context(), &req); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to delete memories", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, gin.H{"deleted": true})
}

func (h *MemoryHandler) Health(ctx *gin.Context) {
	resp, err := h.memory.Health(ctx.Request.Context())
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to check memory health", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, resp)
}
