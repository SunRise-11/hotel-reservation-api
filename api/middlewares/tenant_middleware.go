package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/global_variables"
	"reservation-api/internal/utils"
	"strings"
)

func TenantMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		var tenantStr = strings.TrimSpace(c.Request().Header.Get("X-Tenant-ID"))

		if tenantStr == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "X-Tenant-ID header is null")
		}

		tenantID, _ := utils.ConvertToUint(tenantStr)
		c.Set(global_variables.TenantIDKey, tenantStr)

		ctx := context.WithValue(c.Request().Context(), "TenantID", tenantID)
		c.Set(global_variables.TenantIDCtx, ctx)

		r := c.Request().WithContext(context.WithValue(c.Request().Context(),
			global_variables.CurrentLang, c.Request().Header.Get("Accept-Language")))
		c.SetRequest(r)

		return next(c)
	}
}
