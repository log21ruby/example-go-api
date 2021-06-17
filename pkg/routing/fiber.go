package routing

import (
	"fmt"
	"strings"
	"time"

	"example-go-api/docs"
	"example-go-api/model"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/storage/redis"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	defaultSkipper = func(c *fiber.Ctx) bool {
		return strings.Contains(c.Path(), "swagger")
	}
)

// FiberSkipper skip next
type fiberSkipper func(c *fiber.Ctx) bool

// FiberMiddleware init
type FiberMiddleware struct {
	Skipper fiberSkipper
	Config  interface{}
}

// InitFiber init http fiber
func InitFiber() *FiberMiddleware {
	m := &FiberMiddleware{
		Skipper: defaultSkipper,
		Config: fiber.New(fiber.Config{
			Prefork:       viper.GetBool("app.prefork"),
			CaseSensitive: false,
			StrictRouting: false,
			ErrorHandler: func(f *fiber.Ctx, err error) error {
				code := fiber.StatusInternalServerError
				if e, isOk := err.(*fiber.Error); isOk {
					code = e.Code
				}
				logrus.WithFields(
					logrus.Fields{
						"code": code,
					}).Error(err)
				return fiberError(code, f, err)
			},
		}),
	}

	return m
}

// InitFiberMiddleware fiber use
func (m *FiberMiddleware) InitFiberMiddleware() (*fiber.App, fiber.Router) {
	service := viper.GetString("app.service")
	basePath := "/api/" + service

	f := m.Config.(*fiber.App)
	f.Use(cors.New())
	f.Use(recover.New())
	f.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			if auth := c.Get(fiber.HeaderAuthorization); len(auth) > 0 {
				return true
			}
			if cacheControl := c.Get(fiber.HeaderCacheControl); len(cacheControl) > 0 {
				return true
			}
			return false
		},
		Expiration:   10 * time.Second,
		CacheControl: true,
		Storage: redis.New(redis.Config{
			Host:     viper.GetString("redis.host"),
			Port:     viper.GetInt("redis.port"),
			Password: viper.GetString("redis.password"),
			Database: viper.GetInt("redis.fiberCacheDB"),
			Reset:    false,
		}),
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s?%s", c.Path(), c.Request().URI().QueryArgs().String())
		},
	}))
	f.Use(m.fiberLogResponse)
	f.Use(m.fiberLogRequest)

	router := f.Group(basePath)

	if viper.GetString("app.env") != "prod" {
		docs.SwaggerInfo.Title = "Swagger " + service + " service"
		docs.SwaggerInfo.Host = viper.GetString("app.host")
		docs.SwaggerInfo.BasePath = basePath
		docs.SwaggerInfo.Schemes = []string{"https", "http"}
		router.Get("/swagger/*", swagger.Handler)
	}
	router.Get("/healthcheck", healthcheck)
	return f, router
}

// Healthcheck handler
// ShowAccount godoc
// @Param test query string false "sample query"
// @Success 200 {object} model.BaseResponse{result=bool} "server is ok"
// @Failure default {object} model.BaseErrorResponse{error=int} "server is not ok"
// @Router /healthcheck [get]
func healthcheck(c *fiber.Ctx) error {
	return c.JSON(model.NewBaseResponse(0, "v1.0.0"))
}

func fiberError(code int, f *fiber.Ctx, err error) error {
	switch code {
	case 1:
		return f.Status(fiber.StatusBadRequest).JSON(model.NewBaseErrorResponse(code, err.Error()))
	case 2:
		return f.Status(fiber.StatusBadRequest).JSON(model.NewBaseErrorResponse(code, "bad request"))
	case 3:
		return f.Status(fiber.StatusBadRequest).JSON(model.NewBaseErrorResponse(code, "invalid value"))
	case 4:
		return f.Status(fiber.StatusForbidden).JSON(model.NewBaseErrorResponse(code, "invalid token"))
	default:
		return f.Status(fiber.StatusBadRequest).JSON(model.NewBaseErrorResponse(code, "Internal server error"))
	}
}

func (m *FiberMiddleware) fiberLogRequest(c *fiber.Ctx) error {
	if m.Skipper(c) {
		return c.Next()
	}
	body := string(c.Request().Body())
	method := string(c.Request().Header.Method())

	newSpanID := strings.ReplaceAll(uuid.New().String(), "-", "")
	spanID := c.Get("span_id", newSpanID[len(newSpanID)-20:])

	c.Locals("span_id", spanID)

	logrus.WithFields(
		logrus.Fields{
			"method": method,
			"path":   c.Path(),
			"body":   string(body),
		}).Infof("Request")
	return c.Next()
}

func (m *FiberMiddleware) fiberLogResponse(c *fiber.Ctx) error {
	if m.Skipper(c) {
		return c.Next()
	}
	if err := c.Next(); err != nil {
		return err
	}

	body := string(c.Response().Body())
	method := string(c.Request().Header.Method())
	logrus.WithFields(
		logrus.Fields{
			"method": method,
			"path":   c.Path(),
			"body":   string(body),
		}).Infof("Response")
	return nil
}
