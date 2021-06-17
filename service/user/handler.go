package user

import (
	"example-go-api/model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Handler type
type Handler struct {
	service Servicer
}

// NewHandler init route
func NewHandler(service Servicer) *Handler {
	return &Handler{
		service: service,
	}
}

// GetUserWithCache handle
// @Tags user
// @Param id path int true "user id"
// @Success 200 {object} model.BaseResponse{data=model.User} "server is ok"
// @Failure default {object} model.BaseErrorResponse{code=int,error=string} "server is not ok"
// @Router /v1/user/{id} [post]
func (h *Handler) GetUserWithCache(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, _ := strconv.Atoi(id)
	field := []string{"email_verified", "first_name", "last_name", "first_name_th", "last_name_th", "email", "google_authenticator", "google_authenticator_verified", "trading_credit", "rank_vip", "account_code", "updated_at", "created_at"}
	resp, err := h.service.GetUserWithUIDCache(1*time.Second, "core:user:"+id, uid, field)
	if err != nil {
		return err
	}
	return c.JSON(model.NewBaseResponse(0, resp))
}
