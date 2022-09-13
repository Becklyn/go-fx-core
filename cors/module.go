package cors

import (
	"github.com/Becklyn/go-fx-core/middleware"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
)

var Module = fx.Invoke(
	useCorsMiddleware,
)

func useCorsMiddleware(middleware middleware.Middleware) {
	middleware.Use(cors.New())
}
