package controllers

import (
	"net/http"

	"lsp-api/internal/services"
	"lsp-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	FullName string `json:"full_name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req RegisterRequest

	valid, _ := utils.ValidateRequest(ctx, &req)
	if !valid {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Validation failed"))
		return
	}

	err := c.authService.Register(req.Username, req.FullName, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("User registered successfully", nil))
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest

	valid, _ := utils.ValidateRequest(ctx, &req)
	if !valid {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Validation failed"))
		return
	}

	token, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Login successful", LoginResponse{Token: token}))
}

func (c *AuthController) Logout(ctx *gin.Context) {
	// In a stateless JWT implementation, logout is handled client-side by removing the token
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Logged out successfully", nil))
}

func (c *AuthController) RegisterRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", c.Register)
		authRouter.POST("/login", c.Login)
		authRouter.POST("/logout", authMiddleware, c.Logout)
	}
}
