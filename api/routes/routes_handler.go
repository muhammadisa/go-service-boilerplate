package routes

import (
	"fmt"
	"log"
	"net/http"

	_userApi "github.com/muhammadisa/go-service-boilerplate/api/app/user/delivery/http"
	_userRepo "github.com/muhammadisa/go-service-boilerplate/api/app/user/repository"
	_userUsecase "github.com/muhammadisa/go-service-boilerplate/api/app/user/usecase"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/customvalidator"

	_aliyunOSSApi "github.com/muhammadisa/go-service-boilerplate/api/app/aliyunoss/delivery/http"
	_aliyunOSSRepo "github.com/muhammadisa/go-service-boilerplate/api/app/aliyunoss/repository"
	_aliyunOSSUsecase "github.com/muhammadisa/go-service-boilerplate/api/app/aliyunoss/usecase"

	_foobarApi "github.com/muhammadisa/go-service-boilerplate/api/app/foobar/delivery/http"
	_foobarRepo "github.com/muhammadisa/go-service-boilerplate/api/app/foobar/repository"
	_foobarUsecase "github.com/muhammadisa/go-service-boilerplate/api/app/foobar/usecase"
	"github.com/muhammadisa/go-service-boilerplate/api/cache"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muhammadisa/go-service-boilerplate/api/response"
)

// Routes struct
type Routes struct {
	Echo  *echo.Echo
	Group *echo.Group
	DB    *gorm.DB
	Cache cache.Redis
}

// RouteConfigs struct
type RouteConfigs struct {
	EchoData  *echo.Echo
	DB        *gorm.DB
	APISecret string
	Version   string
	Port      string
	Origins   []string
	Cache     cache.Redis
}

// IRouteConfigs interface
type IRouteConfigs interface {
	NewHTTPRoute()
}

// NewHTTPRoute echo route initialization
func (rc RouteConfigs) NewHTTPRoute() {
	// Initialize route configs
	restful := rc.EchoData.Group(fmt.Sprintf("api/%s", rc.Version))
	handler := &Routes{
		Echo:  rc.EchoData,
		Group: restful,
		DB:    rc.DB,
		Cache: rc.Cache,
	}
	handler.Echo.Validator = &customvalidator.CustomValidator{Validator: validator.New()}
	handler.setupMiddleware(rc.APISecret, rc.Origins)
	handler.setInitRoutes()

	// Internal routers
	handler.initUserRoutes()
	handler.initFoobarRoutes()
	handler.initAliyunOSSRoutes()

	// Starting Echo Server
	log.Fatal(handler.Echo.Start(rc.Port))
}

func (r *Routes) setupMiddleware(apiSecret string, origins []string) {
	// main middleware
	r.Echo.Use(middleware.Recover())
	r.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: origins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodOptions,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
	}))
}

func (r *Routes) setInitRoutes() {
	r.Echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.Response{
			StatusCode: http.StatusOK,
			Message:    "Server is running",
			Data:       true,
		})
	})
}

// Create route initialization function here
func (r *Routes) initUserRoutes() {
	userRepo := _userRepo.NewPostgresUserRepo(r.DB, r.Cache)
	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	_userApi.NewUserHandler(r.Group, userUsecase)
}

func (r *Routes) initFoobarRoutes() {
	foobarRepo := _foobarRepo.NewPostgresFoobarRepo(r.DB, r.Cache)
	foobarUsecase := _foobarUsecase.NewFoobarUsecase(foobarRepo)
	_foobarApi.NewFoobarHandler(r.Group, foobarUsecase)
}

func (r *Routes) initAliyunOSSRoutes() {
	aliyunOSSRepo := _aliyunOSSRepo.NewAliyunOSSInteractorRepo()
	aliyunOSSUsecase := _aliyunOSSUsecase.NewAliyunOSSUsecase(aliyunOSSRepo)
	_aliyunOSSApi.NewAliyunOSSHandler(r.Group, aliyunOSSUsecase)
}
