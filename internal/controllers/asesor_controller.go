package controllers

import (
	"net/http"
	"strconv"

	"lsp-api/internal/services"
	"lsp-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type AsesorController struct {
	asesorService services.AsesorService
}

func NewAsesorController(asesorService services.AsesorService) *AsesorController {
	return &AsesorController{
		asesorService: asesorService,
	}
}

type CreateAsesorRequest struct {
	NamaLengkap  string `json:"nama_lengkap" binding:"required,min=3,max=100"`
	NoRegistrasi string `json:"no_registrasi" binding:"required,min=3,max=50"`
	Email        string `json:"email" binding:"required,email"`
	NoTelepon    string `json:"no_telepon" binding:"required"`
	KompetensiID []uint `json:"kompetensi_id" binding:"required"`
}

type UpdateAsesorRequest struct {
	NamaLengkap  string `json:"nama_lengkap" binding:"required,min=3,max=100"`
	NoRegistrasi string `json:"no_registrasi" binding:"required,min=3,max=50"`
	Email        string `json:"email" binding:"required,email"`
	NoTelepon    string `json:"no_telepon" binding:"required"`
	KompetensiID []uint `json:"kompetensi_id" binding:"required"`
}

func (c *AsesorController) CreateAsesor(ctx *gin.Context) {
	var req CreateAsesorRequest

	valid, errs := utils.ValidateRequest(ctx, &req)
	if !valid {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Validation failed", errs))
		return
	}

	asesor, err := c.asesorService.CreateAsesor(
		req.NamaLengkap,
		req.NoRegistrasi,
		req.Email,
		req.NoTelepon,
		req.KompetensiID,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Asesor created successfully", asesor))
}

func (c *AsesorController) UpdateAsesor(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid asesor ID", nil))
		return
	}

	var req UpdateAsesorRequest

	valid, errs := utils.ValidateRequest(ctx, &req)
	if !valid {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Validation failed", errs))
		return
	}

	asesor, err := c.asesorService.UpdateAsesor(
		uint(id),
		req.NamaLengkap,
		req.NoRegistrasi,
		req.Email,
		req.NoTelepon,
		req.KompetensiID,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Asesor updated successfully", asesor))
}

func (c *AsesorController) DeleteAsesor(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid asesor ID", nil))
		return
	}

	err = c.asesorService.DeleteAsesor(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Asesor deleted successfully", nil))
}

func (c *AsesorController) GetAsesor(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid asesor ID", nil))
		return
	}

	asesor, err := c.asesorService.GetAsesorByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Asesor not found", nil))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Asesor retrieved successfully", asesor))
}

func (c *AsesorController) GetAllAsesors(ctx *gin.Context) {
	asesors, err := c.asesorService.GetAllAsesors()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve asesors", nil))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Asesors retrieved successfully", asesors))
}

func (c *AsesorController) GetAsesorByNoRegistrasi(ctx *gin.Context) {
	noRegistrasi := ctx.Param("no_registrasi")

	asesor, err := c.asesorService.GetAsesorByNoRegistrasi(noRegistrasi)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Asesor not found", nil))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse("Asesor retrieved successfully", asesor))
}

func (c *AsesorController) RegisterRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	asesorRouter := router.Group("/asesors", authMiddleware)
	{
		asesorRouter.POST("/", c.CreateAsesor)
		asesorRouter.PUT("/:id", c.UpdateAsesor)
		asesorRouter.DELETE("/:id", c.DeleteAsesor)
		asesorRouter.GET("/:id", c.GetAsesor)
		asesorRouter.GET("/", c.GetAllAsesors)
		asesorRouter.GET("/registrasi/:no_registrasi", c.GetAsesorByNoRegistrasi)
	}
}
