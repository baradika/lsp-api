package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"lsp-api/internal/services"
	"lsp-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type FormAPL01Controller struct {
	formService    services.FormAPL01Service
	asesmenService services.AsesmenService
}

func NewFormAPL01Controller(formService services.FormAPL01Service, asesmenService services.AsesmenService) *FormAPL01Controller {
	return &FormAPL01Controller{
		formService:    formService,
		asesmenService: asesmenService,
	}
}

func (c *FormAPL01Controller) RegisterRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	apl01Router := router.Group("/apl01", authMiddleware)
	{
		apl01Router.POST("/", c.CreateFormAPL01)
		apl01Router.GET("/", c.GetFormAPL01)
		apl01Router.PUT("/:id", c.UpdateFormAPL01)
		apl01Router.DELETE("/:id", c.DeleteFormAPL01)
	}
}

func (c *FormAPL01Controller) CreateFormAPL01(ctx *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized", nil))
		return
	}

	// Parse request body
	var req map[string]interface{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	// Get asesmen ID from request
	asesmenIDFloat, ok := req["id_asesmen"].(float64)
	if !ok {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Missing or invalid asesmen ID", nil))
		return
	}
	asesmenID := uint(asesmenIDFloat)

	// Verify asesmen belongs to the user
	asesmen, err := c.asesmenService.GetAsesmenByID(asesmenID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Asesmen not found", err))
		return
	}

	// Check if user is the asesi for this asesmen
	if asesmen.Asesi == nil || asesmen.Asesi.UserID != userID.(uint) {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse("You don't have permission to create APL01 for this asesmen", nil))
		return
	}

	// Create form
	form, err := c.formService.CreateFormAPL01(asesmenID, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create form APL01", err))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Form APL01 created successfully", form))
}

func (c *FormAPL01Controller) GetFormAPL01(ctx *gin.Context) {
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

		form, err := c.formService.GetFormAPL01ByID(uint(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Form APL01 not found", err))
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

		ctx.JSON(http.StatusOK, utils.SuccessResponse("Form APL01 retrieved successfully", form))
		return
	}

	// If asesmen_id parameter is provided, get form by asesmen ID
	if asesmenIDStr := ctx.Query("asesmen_id"); asesmenIDStr != "" {
		asesmenID, err := strconv.ParseUint(asesmenIDStr, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid asesmen ID format", err))
			return
		}

		// Check if user has permission to view this asesmen's form
		if role.(string) != "Admin" && role.(string) != "Asesor" {
			asesmen, err := c.asesmenService.GetAsesmenByID(uint(asesmenID))
			if err != nil || asesmen.Asesi == nil || asesmen.Asesi.UserID != userID.(uint) {
				ctx.JSON(http.StatusForbidden, utils.ErrorResponse("You don't have permission to view forms for this asesmen", nil))
				return
			}
		}

		form, err := c.formService.GetFormAPL01ByAsesmenID(uint(asesmenID))
		if err != nil {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Form APL01 not found for this asesmen", err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse("Form APL01 retrieved successfully", form))
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

		forms, err := c.formService.GetFormAPL01ByAsesiID(uint(asesiID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve forms", err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse("Forms APL01 retrieved successfully", forms))
		return
	}

	// If no specific parameter is provided and user is Asesi, get all forms for the user
	if role.(string) == "Asesi" {
		// Get asesi ID for this user
		asesiID, err := c.getAsesiIDFromUserID(userID.(uint))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve asesi information", err))
			return
		}

		forms, err := c.formService.GetFormAPL01ByAsesiID(asesiID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve forms", err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse("Forms APL01 retrieved successfully", forms))
		return
	}

	// For Admin and Asesor, this endpoint should be more specific
	ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Please provide id, asesmen_id, or asesi_id parameter", nil))
}

func (c *FormAPL01Controller) UpdateFormAPL01(ctx *gin.Context) {
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
	existingForm, err := c.formService.GetFormAPL01ByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Form APL01 not found", err))
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
	var req map[string]interface{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body", err))
		return
	}

	// Update form
	updatedForm, err := c.formService.UpdateFormAPL01(uint(id), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update form APL01", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Form APL01 updated successfully", updatedForm))
}

func (c *FormAPL01Controller) DeleteFormAPL01(ctx *gin.Context) {

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized", nil))
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID format", err))
		return
	}

	existingForm, err := c.formService.GetFormAPL01ByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Form APL01 not found", err))
		return
	}

	asesmen, err := c.asesmenService.GetAsesmenByID(*existingForm.IDAsesmen)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Asesmen not found", err))
		return
	}

	if asesmen.Asesi == nil || asesmen.Asesi.UserID != userID.(uint) {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse("You don't have permission to delete this form", nil))
		return
	}

	err = c.formService.DeleteFormAPL01(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete form APL01", err))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Form APL01 deleted successfully", nil))
}

// Fungsi helper untuk mendapatkan ID asesi dari ID user
func (c *FormAPL01Controller) getAsesiIDFromUserID(userID uint) (uint, error) {
	// Dapatkan semua asesmen untuk user ini
	asesmens, err := c.asesmenService.GetAsesmenByAsesiID(userID)
	if err != nil {
		return 0, err
	}

	// Jika ada asesmen, ambil ID asesi dari asesmen pertama
	if len(asesmens) > 0 && asesmens[0].Asesi != nil {
		return *asesmens[0].IDAsesi, nil
	}

	return 0, errors.New("asesi not found for this user")
}
