package routes

import (
	"catering/config"
	"catering/controllers"
	"catering/middleware"
	"catering/models"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	// "github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		hToken := e.Request().Header.Get("Authorization")

		var tokenString string
		arrayToken := strings.Split(hToken, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := middleware.ValidateToken(tokenString)

		if err != nil {
			return e.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			return e.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
		}

		userId := int(claim["user_id"].(float64))
		user := models.User{}
		errors := config.DB.Where("id=?", userId).First(&userId).Error

		if errors != nil {
			return e.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
		}

		e.Set("currentUser", user)

		return next(e)
	}
}

func Init() *echo.Echo {
	e := echo.New()
	route := e.Group("/")

	route.POST("image", controllers.CreateImage)
	route.DELETE("image/:id", controllers.DeleteImage)

	route.POST("category", controllers.CreateCategory)
	route.GET("categories", controllers.GetCategories)
	route.GET("category/:id", controllers.ShowCategory)
	route.PUT("category/:id", controllers.UpdateCategory)
	route.DELETE("category/:id", controllers.DeleteCategory)

	route.POST("regency", controllers.CreateRegency)
	route.GET("regencies", controllers.GetRegency)
	route.GET("regency/:id", controllers.ShowRegency)
	route.PUT("regency/:id", controllers.UpdateRegency)
	route.DELETE("regency/:id", controllers.DeleteRegency)

	route.POST("paket", controllers.CreatePaket)
	route.GET("pakets", controllers.GetPakets)
	route.GET("paket/:id", controllers.ShowPaket)
	route.PUT("paket/:id", controllers.UpdatePaket)
	route.DELETE("paket/:id", controllers.DeletePaket)

	route.POST("user", controllers.CreateUser)
	route.GET("users", controllers.GetUsers)
	route.GET("user/:id", controllers.ShowUser)
	route.PUT("user/:id", controllers.UpdateUser)
	route.DELETE("user/:id", controllers.DeleteUser)

	route.POST("transaction", controllers.CreateTransaction)
	route.GET("transactions", controllers.GetTransaction)
	route.GET("transactions/user", controllers.GetTransactionByUserId)
	route.POST("transaction/notif", controllers.GetNotification)

	route.POST("registrasi", controllers.Registration)
	route.POST("login", controllers.Login)

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
