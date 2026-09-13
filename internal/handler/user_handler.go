package handler

import (
	"net/http"
	"safrenz-go-boilerplate/internal/service"
	"safrenz-go-boilerplate/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"SFR-01-003",
			"Gagal mengambil data user",
		)
	}

	return utils.SuccessResponse(
		c,
		http.StatusOK,
		"SFR-01-001",
		"Berhasil mengambil data user",
		users,
	)
}
