package posts

import (
	"github.com/go-jedi/gateway/internal/domain/posts"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) create(c fiber.Ctx) error {
	h.logger.Debug("[create a new post] execute handler")

	var dto posts.CreateDTO
	if err := c.Bind().Body(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.postsService.Create(c.Context(), dto)
	if err != nil {
		h.logger.Error("failed to create post", "error", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(result)
}
