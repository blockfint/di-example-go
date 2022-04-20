package handler

import (
	"net/http"
	"strconv"

	"github.com/blockfint/di-example-go/app/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) NewCustomersGroup() *group {
	return &group{
		Prefix:      "/customers",
		middlewares: []echo.MiddlewareFunc{},
		routes: []route{
			{
				Path:        "",
				Method:      http.MethodGet,
				HandlerFunc: h.ListCustomers,
			},
			{
				Path:        "/:customerID",
				Method:      http.MethodGet,
				HandlerFunc: h.FindCustomerByID,
			},
		},
	}
}

func (h *Handler) ListCustomers(c echo.Context) error {
	customers, err := h.customerController.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, customers)
}

func (h *Handler) FindCustomerByID(c echo.Context) error {
	id64, err := strconv.ParseUint(c.Param("customerID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	customerID := uint(id64)
	customer, err := h.customerController.FindByID(customerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if customer == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	return c.JSON(http.StatusOK, customer)
}
