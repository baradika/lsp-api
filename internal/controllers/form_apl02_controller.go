package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"lsp-api/internal/services"
	"lsp-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type FormAPL02Controller struct {
	formService    services.FormAPL02Service
	asesmenService services.AsesmenService
}

func NewFormAPL02Controller(formService services.FormAPL02Service, asesmenService services.AsesmenService) *FormAPL02Controller {
	return &FormAPL02Controller{
		formService:    formService,
		asesmenService: asesmenService,
	}
}

func (c *FormAPL02Controller) RegisterRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	apl02Router := router.Group("/apl02", authMiddleware)
	{
		apl02Router.POST("/", c.CreateFormAPL02)
		apl02Router.GET("/", c.GetFormAPL02)
		apl02Router.PUT("/:id", c.UpdateFormAPL02)
		apl02Router.DELETE("/:id", c.DeleteFormAPL02)
	}
}

func (c *FormAPL02Controller) CreateFormAPL02(ctx *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized", nil))
		return
	}

	// Parse request body
	var req struct {
		IDAsesmen uint   `json:"id_asesmen" binding:"required"`
		KodeUnit  string `json:"kode_unit" binding:"required"`
		JudulUnit string `json:"judul_unit" binding:"required"`
		Elemen    string `json:"elemen" binding:"required"`
		KUK       string `json:"kuk" binding:"required"`
		Status    string `json:"status" binding:"required,oneof=K BK"`
		Bukti     string `json:"bukti" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	// Verify asesmen belongs to the user
	asesmen, err := c.asesmenService.GetAsesmenByID(req.IDAsesmen)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Asesmen not found", err))
		return
	}

	// Check if user is the asesi for this asesmen
	if asesmen.Asesi == nil || asesmen.Asesi.UserID != userID.(uint) {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse("You don't have permission to create APL02 for this asesmen", nil))
		return
	}

	// Create form
	form, err := c.formService.CreateFormAPL02(
		req.IDAsesmen,
		req.KodeUnit,
		req.JudulUnit,
		req.Elemen,
		req.KUK,
		req.Status,
		req.Bukti,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create form APL02", err))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Form APL02 created successfully", form))
}

func (c *FormAPL02Controller) GetFormAPL02(ctx *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized", nil))
		return
	}

	// Get user role from context
	role, exists := ctx.Get("role")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized", nil))
		return
	}

	// If ID parameter is provided, get specific form
	if idStr := ctx.Query("id"); idStr != "" {
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID format", err))
			return
		}

		form, err := c.formService.GetFormAPL02ByID(uint(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Form APL02 not found", err))
			return
		}

		// Check if user has permission to view this form
		if role.(string) != "Admin" && role.(string) != "Asesor" {
			// For Asesi, check if the form belongs to them
			asesmen, err := c.asesmenService.GetAsesmenByID(*form.IDAsesmen)
			if err != nil || asesmen.Asesi == nil || asesmen.Asesi.UserID != userID.(uint) {
				ctx.JSON(http.StatusForbidden, utils.ErrorResponse("You don't have permission to view this form", nil))
				return
			}
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse("Form APL02 retrieved successfully", form))
		return
	}

	// If asesmen_id parameter is provided, get forms by asesmen ID
	if asesmenIDStr := ctx.Query("asesmen_id"); asesmenIDStr != "" {
		asesmenID, err := strconv.ParseUint(asesmenIDStr, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid asesmen ID format", err))
			return
		}

		// Check if user has permission to view this asesmen's forms
		if role.(string) != "Admin" && role.(string) != "Asesor" {
			asesmen, err := c.asesmenService.GetAsesmenByID(uint(asesmenID))
			if err != nil || asesmen.Asesi == nil || asesmen.Asesi.UserID != userID.(uint) {
				ctx.JSON(http.StatusForbidden, utils.ErrorResponse("You don't have permission to view forms for this asesmen", nil))
				return
			}
		}

		forms, err := c.formService.GetFormAPL02ByAsesmenID(uint(asesmenID))
		if err != nil {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Forms APL02 not found for this asesmen", err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse("Forms APL02 retrieved successfully", forms))
		return
	}

	// If asesi_id parameter is provided, get forms by asesi ID
	if asesiIDStr := ctx.Query("asesi_id"); asesiIDStr != "" {
		asesiID, err := strconv.ParseUint(asesiIDStr, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid asesi ID format", err))
			return
		}

		// Check if user has permission to view this asesi's forms
		if role.(string) == "Asesi" && userID.(uint) != uint(asesiID) {
			ctx.JSON(http.StatusForbidden, utils.ErrorResponse("You don't have permission to view forms for this asesi", nil))
			return
		}

		forms, err := c.formService.GetFormAPL02ByAsesiID(uint(asesiID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve forms", err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse("Forms APL02 retrieved successfully", forms))
		return
	}

	// For Admin and Asesor, this endpoint should be more specific
	ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Please provide id, asesmen_id, or asesi_id parameter", nil))
}

func (c *FormAPL02Controller) UpdateFormAPL02(ctx *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized", nil))
		return
	}

	// Get form ID from URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID format", err))
		return
	}

	// Get existing form
	existingForm, err := c.formService.GetFormAPL02ByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Form APL02 not found", err))
		return
	}

	// Verify asesmen belongs to the user
	asesmen, err := c.asesmenService.GetAsesmenByID(*existingForm.IDAsesmen)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Asesmen not found", err))
		return
	}

	// Check if user is the asesi for this asesmen
	if asesmen.Asesi == nil || asesmen.Asesi.UserID != userID.(uint) {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse("You don't have permission to update this form", nil))
		return
	}

	// Parse request body
	var req struct {
		KodeUnit  string `json:"kode_unit"`
		JudulUnit string `json:"judul_unit"`
		Elemen    string `json:"elemen"`
		KUK       string `json:"kuk"`
		Status    string `json:"status" binding:"omitempty,oneof=K BK"`
		Bukti     string `json:"bukti"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	// Update form
	updatedForm, err := c.formService.UpdateFormAPL02(
		uint(id),
		req.KodeUnit,
		req.JudulUnit,
		req.Elemen,
		req.KUK,
		req.Status,
		req.Bukti,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update form APL02", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Form APL02 updated successfully", updatedForm))
}

func (c *FormAPL02Controller) DeleteFormAPL02(ctx *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized", nil))
		return
	}

	// Get form ID from URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID format", err))
		return
	}

	// Get existing form
	existingForm, err := c.formService.GetFormAPL02ByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Form APL02 not found", err))
		return
	}

	// Verify asesmen belongs to the user
	asesmen, err := c.asesmenService.GetAsesmenByID(*existingForm.IDAsesmen)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Asesmen not found", err))
		return
	}

	// Check if user is the asesi for this asesmen
	if asesmen.Asesi == nil || asesmen.Asesi.UserID != userID.(uint) {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse("You don't have permission to delete this form", nil))
		return
	}

	// Delete form
	err = c.formService.DeleteFormAPL02(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete form APL02", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Form APL02 deleted successfully", nil))
}

// Helper function to get asesi ID from user ID
func (c *FormAPL02Controller) getAsesiIDFromUserID(userID uint) (uint, error) {
	// Get all asesmen for this user
	asesmens, err := c.asesmenService.GetAsesmenByAsesiID(userID)
	if err != nil {
		return 0, err
	}

	if len(asesmens) == 0 || asesmens[0].Asesi == nil {
		return 0, errors.New("no asesmen found for this user")
	}

	return asesmens[0].Asesi.ID, nil
}
