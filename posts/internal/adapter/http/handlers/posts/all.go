package posts

import (
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) all(c fiber.Ctx) error {
	h.logger.Debug("[get all posts] execute handler")

	result, err := h.postsService.All(c.Context())
	if err != nil {
		h.logger.Error("failed to get all posts", "error", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(result)
}
