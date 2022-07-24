package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/handler/request"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/respond"
)

func (h *HandlerV1) SaveDeposit() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		service := h.serviceRegistry.GetAccountService()
		var req request.DepositRequestParams
		if err := c.QueryParser(&req); err != nil {
			return respond.Fail(c, http.StatusBadRequest, http.StatusBadRequest, err)
		}

		err := service.Deposit(req.WalletID, req.Amount)
		if err != nil {
			return respond.Fail(c, http.StatusInternalServerError, http.StatusInternalServerError, err)
		}

		return respond.Success(c, http.StatusAccepted, req)
	}
}

func (h *HandlerV1) GetBalance() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		service := h.serviceRegistry.GetAccountService()
		var req request.BalanceRequestParams
		if err := c.QueryParser(&req); err != nil {
			return respond.Fail(c, http.StatusBadRequest, http.StatusBadRequest, err)
		}

		amount, err := service.Balance(req.WalletID)
		if err != nil {
			return respond.Fail(c, http.StatusInternalServerError, http.StatusInternalServerError, err)
		}

		return respond.Success(c, http.StatusOK, amount)
	}
}
