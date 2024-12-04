package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hcd233/Aris-url-gen/internal/api/dto"
	"github.com/hcd233/Aris-url-gen/internal/api/service"
)

type shortenUrlHandler struct {
	service service.ShortURLService
}

type ShortURLHandlerOption func(handler *shortenUrlHandler)

func WithShortURLService(service service.ShortURLService) ShortURLHandlerOption {
	return func(handler *shortenUrlHandler) {
		handler.service = service
	}
}

func NewShortURLHandler(handlerOptions ...ShortURLHandlerOption) ShortURLHandler {
	handler := &shortenUrlHandler{}
	for _, option := range handlerOptions {
		option(handler)
	}
	return handler
}

func (h *shortenUrlHandler) GenerateShortURL(ctx *fiber.Ctx) error {
	request := new(dto.GenerateShortUrlRequest)
	if err := ctx.BodyParser(request); err != nil {
		code := dto.CodeInvalidRequest
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.StandardResponse{
			Code:    code,
			Message: code.ToMessage(err.Error()),
		})
	}

	response, err := h.service.GenerateShortURL(request)
	if err != nil {
		code := dto.CodeGenerateShortURLFailed
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.StandardResponse{
			Code:    code,
			Message: code.ToMessage(err.Error()),
		})
	}

	return ctx.JSON(dto.StandardResponse{
		Code:    dto.CodeOK,
		Message: dto.CodeOK.ToMessage(),
		Data:    response,
	})
}

func (h *shortenUrlHandler) GetOriginalURL(ctx *fiber.Ctx) error {
	request := new(dto.GetOriginalUrlRequest)
	if err := ctx.ParamsParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response, err := h.service.GetOriginalURL(request)
	if err != nil {
		code := dto.CodeGetOriginalURLFailed
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.StandardResponse{
			Code:    code,
			Message: code.ToMessage(err.Error()),
		})
	}

	return ctx.Redirect(response.OriginalURL)
}
