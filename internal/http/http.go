package http

import (
	"context"
	"encoding/json"
	"github.com/Nicholas2012/task_tz/internal/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Users interface {
	Random(ctx context.Context) (storage.User, error)
	List(ctx context.Context) ([]storage.User, error)
}

type Handlers struct {
	users Users
}

func New(users Users) *Handlers {
	return &Handlers{
		users: users,
	}
}

func (h *Handlers) Register(r *gin.Engine) {
	r.GET("/random", h.Random)
	r.GET("/list", h.List)
}

func (h *Handlers) Random(ctx *gin.Context) {
	user, err := h.users.Random(ctx.Request.Context())
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	h.ok(ctx, user)
}

func (h *Handlers) List(ctx *gin.Context) {
	users, err := h.users.List(ctx.Request.Context())
	if err != nil {
		h.handleError(ctx, err)
		return
	}

	h.ok(ctx, users)
}

func (h *Handlers) ok(ctx *gin.Context, data any) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(ctx.Writer).Encode(data); err != nil {
		h.handleError(ctx, err)
		return
	}
}

func (h *Handlers) handleError(ctx *gin.Context, err error) {
	log.Printf("[ERROR] error durning request, request %s, error %v", ctx.Request.URL.String(), err)
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
