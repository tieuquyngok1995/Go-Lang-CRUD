package auth

import (
	"go-crud/internal/middleware"
	"go-crud/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	// Public
	auth.POST("/login", h.Login)

	// Protected
	auth.Use(middleware.Auth())
	auth.POST("/logout", h.Logout)
	auth.POST("/logout-all", h.LogoutAll)
}

// POST /api/v1/auth/login
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(c, http.StatusOK, res)
}

// POST /api/v1/auth/logout  ? Header: Authorization: Bearer <token>
func (h *Handler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if err := h.svc.Logout(c.Request.Context(), token); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, gin.H{"message": "logged out successfully"})
}

// POST /api/v1/auth/logout-all  ? Header: Authorization: Bearer <token>
func (h *Handler) LogoutAll(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.Error(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := h.svc.LogoutAll(c.Request.Context(), userID); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, gin.H{"message": "all sessions logged out"})
}
