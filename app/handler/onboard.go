package handler

import (
	"net/http"

	"github.com/blockfint/di-example-go/app/db/model"
	"github.com/blockfint/di-example-go/app/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) NewOnboardGroup() *group {
	return &group{
		Prefix:      "/onboard",
		middlewares: []echo.MiddlewareFunc{},
		routes: []route{
			{
				Path:        "/callback",
				Method:      http.MethodPost,
				HandlerFunc: h.HandleOnboardCallback,
			},
		},
	}
}

func (h *Handler) HandleOnboardCallback(c echo.Context) error {
	var customer model.Customer

	req := &onboardCallbackRequest{}
	if err := req.bind(c, &customer); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := h.customerController.Create(&customer); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.NoContent(http.StatusCreated)
}
