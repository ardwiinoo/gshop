package transaction

import (
	"fmt"
	"net/http"

	infrafiber "github.com/ardwiinoo/online-shop/infra/fiber"
	"github.com/ardwiinoo/online-shop/infra/response"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) CreateTransaction(ctx *fiber.Ctx) error {
	var req = CreateTransactionRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(err),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	req.UserPublicId = fmt.Sprintf("%v", ctx.Locals("PUBLIC_ID"))

	if err := h.svc.CreateTransaction(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]

		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("create transaction success"),
	).Send(ctx)
}

func (h handler) GetTransactionByUser(ctx *fiber.Ctx) error {
	UserPublicId := fmt.Sprintf("%v", ctx.Locals("PUBLIC_ID"))

	trxs, err := h.svc.TransactionHistories(ctx.UserContext(), UserPublicId)

	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]

		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	var response = []TransactionHistoryResponse{}

	for _, trx := range trxs {
		response = append(response, trx.ToTransactionHistoryReponse())
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get transaction histories success"),
		infrafiber.WithPayload(response),
	).Send(ctx)
}