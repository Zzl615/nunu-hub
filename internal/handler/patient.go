package handler

import (
	v1 "feature/api/v1"
	"feature/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PatientHandler interface {
	GetPatientById(ctx *gin.Context)
}

func NewPatientHandler(handler *Handler, patientService service.PatientService) PatientHandler {
	return &patientHandler{
		Handler:        handler,
		patientService: patientService,
	}
}

type patientHandler struct {
	*Handler
	patientService service.PatientService
}

type GetPatientParams struct {
	Id int64 `form:"id" binding:"required"`
}

// Login godoc
// @Summary 获取患者信息
// @Schemes
// @Description
// @Tags 患者模块
// @Accept json
// @Produce json
// @Param request body GetPatientParams true "params"
// @Success 200 {object} v1.LoginResponse
// @Router /ai_patient [post]
func (h *patientHandler) GetPatientById(ctx *gin.Context) {
	var params GetPatientParams
	if err := ctx.ShouldBind(&params); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	patient, err := h.patientService.GetPatientById(params.Id)
	h.logger.Info("GetPatientByID", zap.Any("patient", patient))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}
	v1.HandleSuccess(ctx, patient)
}
