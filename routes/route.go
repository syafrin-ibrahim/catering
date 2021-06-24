package routes

import (
	"catering/controllers"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	// "github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	auth := e.Group("/")
	auth.POST("regency", controllers.CreateRegency)
	auth.GET("regency", controllers.GetRegency)
	auth.POST("paket", controllers.CreatePaket)
	auth.GET("paket", controllers.GetPakets)
	auth.POST("user", controllers.CreateUser)
	auth.GET("user", controllers.GetUsers)
	auth.POST("transaction", controllers.CreateTransaction)
	// auth.POST("registrasi", controllers.Registrasi)

	// admin := e.Group("/")
	// admin.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey:  []byte(constant.SECRET_JWT),
	// 	TokenLookup: "header:authorization",
	// }))

	// //role staff and admin
	// admin.GET("cars", controllers.GetCars)
	// admin.POST("order", controllers.CreateOrder)
	// admin.POST("order/done", controllers.OrderDone)
	// admin.GET("order", controllers.GetOrder)

	// //role admin
	// admin.POST("customer", controllers.CreateCustomer, AdminMiddleware)
	// admin.GET("customers", controllers.GetCustomers, AdminMiddleware)
	// admin.GET("customer/:id", controllers.ShowCustomer, AdminMiddleware)
	// admin.PUT("customer/:id", controllers.UpdateCustomer, AdminMiddleware)
	// admin.DELETE("customer/:id", controllers.DeleteCustomer, AdminMiddleware)

	// admin.POST("car", controllers.CreateCar, AdminMiddleware)
	// admin.GET("car/:id", controllers.ShowCar, AdminMiddleware)
	// admin.PUT("car/:id", controllers.UpdateCar, AdminMiddleware)
	// admin.DELETE("car/:id", controllers.DeleteCar, AdminMiddleware)

	// admin.POST("garage", controllers.CreateGarage, AdminMiddleware)
	// admin.PUT("garage/:id", controllers.UpdateGarage, AdminMiddleware)
	// admin.GET("garage/:id", controllers.ShowGarage, AdminMiddleware)
	// admin.GET("garages", controllers.GetGarages, AdminMiddleware)
	// admin.DELETE("garage/:id", controllers.DeleteGarage, AdminMiddleware)

	return e
}
