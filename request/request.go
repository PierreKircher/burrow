package request

import (
	"github.com/carrot/burrow/controllers"
	"github.com/carrot/burrow/middleware"
	"github.com/carrot/burrow/response"
	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
)

func BuildEcho() *echo.Echo {
	// ----------
	// Framework
	// ----------

	e := echo.New()

	// -----------
	// Middleware
	// -----------

	e.Use(echo_middleware.Logger())
	e.Use(middleware.Recover())

	// -------------------
	// HTTP Error Handler
	// -------------------

	e.SetHTTPErrorHandler(func(err error, context *echo.Context) {
		httpError, ok := err.(*echo.HTTPError)
		if ok {
			response := response.New(context)
			response.SetResponse(httpError.Code(), nil)
			response.Render()
		}
	})

	// ------------
	// Controllers
	// ------------

	topicsController := new(controllers.TopicsController)

	// ----------
	// Endpoints
	// ----------

	e.Get("/topics", topicsController.Index)
	e.Get("/topics/:id", topicsController.Show)
	e.Post("/topics", topicsController.Create)
	e.Put("/topics/:id", topicsController.Update)
	e.Delete("/topics/:id", topicsController.Delete)

	return e
}
