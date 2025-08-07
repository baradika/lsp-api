package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lsp-api/internal/services"
	"lsp-api/internal/utils"
)

type AuthController struct {
	authService services.AuthService
	db          *gorm.DB
}

func (c *AuthController) RegisterRoutes(apiV1 *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	apiV1.POST("/auth/register", c.Register)
	apiV1.POST("/auth/login/admin", c.AdminLogin)
	apiV1.POST("/auth/login/asesor", c.AsesorLogin)
	apiV1.POST("/auth/login/asesi", c.AsesiLogin)
	apiV1.POST("/auth/logout", authMiddleware, c.Logout)
}

func NewAuthController(authService services.AuthService, db *gorm.DB) *AuthController {
	return &AuthController{
		authService: authService,
		db:          db,
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Ganti RegisterAsesi menjadi Register
func (c *AuthController) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	if err := c.authService.Register(req.Username, req.Email, req.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to register user", err))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("User registered successfully", nil))
}

func (c *AuthController) Logout(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Logged out successfully", nil))
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (c *AuthController) AdminLogin(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	token, err := c.authService.AdminLogin(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid credentials", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Login successful", gin.H{
		"token": token,
	}))
}

func (c *AuthController) AsesorLogin(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	token, err := c.authService.AsesorLogin(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid credentials", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Login successful", gin.H{
		"token": token,
	}))
}

func (c *AuthController) AsesiLogin(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	token, err := c.authService.AsesiLogin(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid credentials", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Login successful", gin.H{
		"token": token,
	}))
}
