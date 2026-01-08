package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xerdin442/api-practice/internal/api/dto"
)

func (h *RouteHandler) Signup(c *gin.Context) {
	var req dto.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *RouteHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := h.services.User.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *RouteHandler) Logout(c *gin.Context) {}

func (h *RouteHandler) GetProfile(c *gin.Context) {}

func (h *RouteHandler) UpdateProfile(c *gin.Context) {}
