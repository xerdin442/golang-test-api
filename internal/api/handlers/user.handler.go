package handlers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/xerdin442/api-practice/internal/api/dto"
	"github.com/xerdin442/api-practice/internal/util"
)

func (h *RouteHandler) Signup(c *gin.Context) {
	logger := log.Ctx(c.Request.Context())

	var req dto.SignupRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error().Err(err).Msg("Signup error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	user, err := h.services.User.Signup(c.Request.Context(), req, h.tasksQueue)
	if err != nil {
		logger.Error().Err(err).Msg("Signup error")

		switch {
		case errors.Is(err, util.ErrEmailAlreadyExists), errors.Is(err, util.ErrFileUploadFailed):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		case errors.Is(err, util.ErrUnsupportedImageType):
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": err.Error()})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured during signup"})
		}

		return
	}

	logger.Info().Str("email", req.Email).Msg("Signup successful")
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *RouteHandler) Login(c *gin.Context) {
	logger := log.Ctx(c.Request.Context())

	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error().Err(err).Msg("Login error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	token, err := h.services.User.Login(c.Request.Context(), req)
	if err != nil {
		logger.Error().Err(err).Msg("Login error")

		switch {
		case errors.Is(err, util.ErrInvalidEmail), errors.Is(err, util.ErrInvalidPassword):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured during login"})
		}

		return
	}

	logger.Info().Str("email", req.Email).Msg("Login successful")
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *RouteHandler) Logout(c *gin.Context) {
	logger := log.Ctx(c.Request.Context())

	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	exp, _ := c.Get("token_exp")
	tokenExp := exp.(time.Time)

	err := h.cache.SetJTI(c.Request.Context(), tokenString, "blacklisted", tokenExp)
	if err != nil {
		logger.Error().Err(err).Msg("Logout error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token blacklist error"})

		return
	}

	logger.Info().Msg("Logout successful")
	c.JSON(200, gin.H{"message": "Logged out!"})
}

func (h *RouteHandler) GetProfile(c *gin.Context) {
	logger := log.Ctx(c.Request.Context())

	userID := c.MustGet("userID").(int32)

	user, err := h.services.User.GetProfile(c.Request.Context(), userID)
	if err != nil {
		logger.Error().Err(err).Msg("User profile fetch error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user profile"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
