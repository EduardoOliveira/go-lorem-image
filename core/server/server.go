package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"github.com/eduardooliveira/go-lorem-image/core/config"
)

var e *echo.Echo

func init() {
	e = echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${remote_ip} - - ${time_rfc3339} ${latency_human} "${method} ${uri}" ${status} ${bytes_out} "${refer}" "${user_agent}"` + "\n",
	}))

	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS, echo.HEAD},
		AllowHeaders: []string{"Content-Type", "Authorization", "Accept", "Content-Type", "Origin"},
	}))
	// no route defaults
	e.GET("/", statusOk)
	e.GET("/.health-check", statusOk)
	e.GET("/favicon.ico", statusOk)

}

func GetGroup(path string) *echo.Group {
	return e.Group(path)
}

func Start() {
	e.Start(config.Config().GetString("api.listen"))
}

func statusOk(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
