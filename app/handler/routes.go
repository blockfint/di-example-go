package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(g *echo.Group) {
	routesGroups := []group{
		*h.NewTodosGroup(),
		*h.NewCustomersGroup(),
		*h.NewOnboardGroup(),
	}

	for _, routesGroup := range routesGroups {
		r := g.Group(routesGroup.Prefix, routesGroup.middlewares...)

		for _, route := range routesGroup.routes {
			switch route.Method {
			case http.MethodGet:
				r.GET(route.Path, route.HandlerFunc)
			case http.MethodPost:
				r.POST(route.Path, route.HandlerFunc)
			case http.MethodPut:
				r.PUT(route.Path, route.HandlerFunc)
			case http.MethodPatch:
				r.PATCH(route.Path, route.HandlerFunc)
			default:
				h.logger.Errorf("Unhandle http method: %s", route.Method)
			}
		}
	}
}
