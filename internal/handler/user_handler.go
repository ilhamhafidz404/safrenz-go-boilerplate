package handler

import (
	"net/http"
	"strconv"

	"safrenz-go-boilerplate/internal/dto"
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

// GetUsers godoc
// @Summary      Ambil semua data user
// @Description  Mengambil daftar seluruh user beserta nama role dari third-party API/Redis
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "Berhasil mengambil data user"
// @Failure      500  {object}  map[string]interface{}  "Gagal mengambil data user"
// @Router       /api/v1/users [get]
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers(c.UserContext())
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "SFR-01-003", "Gagal mengambil data user")
	}
	return utils.SuccessResponse(c, http.StatusOK, "SFR-01-001", "Berhasil mengambil data user", users)
}

// GetUserByID godoc
// @Summary      Ambil detail user berdasarkan ID
// @Description  Mengambil detail informasi user berdasarkan ID yang diberikan
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]interface{}  "Berhasil mengambil detail user"
// @Failure      400  {object}  map[string]interface{}  "ID user tidak valid"
// @Failure      404  {object}  map[string]interface{}  "User tidak ditemukan"
// @Router       /api/v1/users/{id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "SFR-01-003", "ID user tidak valid")
	}

	user, err := h.userService.GetUserByID(c.UserContext(), uint(id))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "SFR-01-001", err.Error())
	}
	return utils.SuccessResponse(c, http.StatusOK, "SFR-01-001", "Berhasil mengambil detail user", user)
}

// CreateUser godoc
// @Summary      Buat user baru
// @Description  Menambahkan data user baru ke dalam database
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateUserRequest  true  "Payload pembuatan user"
// @Success      201      {object}  map[string]interface{} "Berhasil membuat user baru"
// @Failure      400      {object}  map[string]interface{} "Payload request tidak valid"
// @Failure      500      {object}  map[string]interface{} "Gagal membuat user baru"
// @Router       /api/v1/users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "SFR-01-003", "Payload request tidak valid")
	}

	user, err := h.userService.CreateUser(c.UserContext(), req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "SFR-01-002", "Gagal membuat user baru")
	}
	return utils.SuccessResponse(c, http.StatusCreated, "SFR-01-001", "Berhasil membuat user baru", user)
}

// UpdateUser godoc
// @Summary      Perbarui data user
// @Description  Perbarui data user yang sudah ada berdasarkan ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id       path      int                    true  "User ID"
// @Param        request  body      dto.UpdateUserRequest  true  "Payload perbaikan user"
// @Success      200      {object}  map[string]interface{} "Berhasil memperbarui data user"
// @Failure      400      {object}  map[string]interface{} "ID atau payload request tidak valid"
// @Failure      500      {object}  map[string]interface{} "Gagal memperbarui user"
// @Router       /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "SFR-01-003", "ID user tidak valid")
	}

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "SFR-01-003", "Payload request tidak valid")
	}

	user, err := h.userService.UpdateUser(c.UserContext(), uint(id), req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "SFR-01-002", "Gagal memperbarui user")
	}
	return utils.SuccessResponse(c, http.StatusOK, "SFR-01-001", "Berhasil memperbarui data user", user)
}

// DeleteUser godoc
// @Summary      Hapus data user
// @Description  Menghapus user dari database berdasarkan ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]interface{}  "Berhasil menghapus user"
// @Failure      400  {object}  map[string]interface{}  "ID user tidak valid"
// @Failure      404  {object}  map[string]interface{}  "User tidak ditemukan"
// @Router       /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "SFR-01-003", "ID user tidak valid")
	}

	if err := h.userService.DeleteUser(c.UserContext(), uint(id)); err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "SFR-01-002", err.Error())
	}
	return utils.SuccessResponse(c, http.StatusOK, "SFR-01-001", "Berhasil menghapus user", nil)
}
