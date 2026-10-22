package main

import (
	"log"
	"management-warehouse/auth"
	"management-warehouse/category"
	"management-warehouse/config"
	"management-warehouse/handler"
	"management-warehouse/middleware"
	"management-warehouse/mutation"
	"management-warehouse/product"
	"management-warehouse/stockadjustment"
	"management-warehouse/supplier"
	"management-warehouse/user"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		origin := c.Request.Header.Get("Origin")

// 		// allow your frontend
// 		if origin == "http://localhost:5173" {
// 			c.Header("Access-Control-Allow-Origin", origin)
// 			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
// 			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 			c.Header("Access-Control-Allow-Credentials", "true")
// 		}

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusNoContent)
// 			return
// 		}

// 		c.Next()
// 	}
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	userRepository := user.NewRepository(config.DB)
	supplierRepository := supplier.NewRepository(config.DB)
	categoryRepository := category.NewRepository(config.DB)
	productRepository := product.NewRepository(config.DB)
	mutationRepository := mutation.NewRepository(config.DB)
	stockAdjustmentRepository := stockadjustment.NewRepository(config.DB)

	userService := user.NewService(userRepository)
	supplierService := supplier.NewService(supplierRepository)
	categoryService := category.NewService(categoryRepository)
	productService := product.NewService(productRepository)
	mutationService := mutation.NewService(mutationRepository)
	stockAdjustmentService := stockadjustment.NewService(stockAdjustmentRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	supplierHandler := handler.NewSupplierHandler(supplierService, authService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	ProductHandler := handler.NewProductHandler(productService) //make protect by auth
	MutationHandler := handler.NewMutationHandler(mutationService)
	stockAdjustmentHandler := handler.NewStockAdjustmentHandler(stockAdjustmentService)

	r := gin.Default()
	r.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(os.Getenv(("SECRET_KEY"))))
	r.Use(sessions.Sessions("warehouseMag", cookieStore))

	api := r.Group("/api/v1")
	//route lists

	api.POST("/auth/register", userHandler.RegisterUser)
	api.POST("/auth/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)

	//user API
	api.GET("/user/detail/:id", middleware.AuthMiddleware(authService, userService), userHandler.Edit)
	api.PUT("/user/update/:id", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin"), userHandler.UpdateUser)
	api.GET("/user/data/:id", middleware.AuthMiddleware(authService, userService), userHandler.DetailUser)

	//supplier API
	api.POST("/supplier/create", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin"), supplierHandler.RegisterSupplier)
	api.GET("/suppliers", middleware.AuthMiddleware(authService, userService), supplierHandler.GetAllSupplier)
	api.GET("/supplier/:id", middleware.AuthMiddleware(authService, userService), supplierHandler.GetSupplierByID)
	api.PUT("/supplier/:id", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin"), supplierHandler.UpdateSupplier)
	api.DELETE("/supplier/:id", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin"), supplierHandler.DeleteSupplier)

	//category API
	api.GET("/categories", middleware.AuthMiddleware(authService, userService), categoryHandler.GetAll)
	api.GET("/category/:id", middleware.AuthMiddleware(authService, userService), categoryHandler.GetByID)
	api.PATCH("/category/:id", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin"), categoryHandler.UpdateCategory)
	api.POST("/category", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin"), categoryHandler.CreateCategory)

	//product API
	api.POST("/product/create", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin"), ProductHandler.CreateProduct)
	api.GET("/product/supplier/:supplier_id", middleware.AuthMiddleware(authService, userService), ProductHandler.GetSupplierProducts)
	api.GET("/products", middleware.AuthMiddleware(authService, userService), ProductHandler.GetAllProduct)
	api.GET("/product/:id", middleware.AuthMiddleware(authService, userService), ProductHandler.GetProductByID)
	api.PUT("/product/:id", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin"), ProductHandler.UpdateProduct)
	api.DELETE("/product/:id", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin"), ProductHandler.DeleteProduct)

	//mutation API
	api.POST("/mutation", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin,stocker"), MutationHandler.CreateMutation)
	api.GET("/mutation", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin,stocker"), MutationHandler.GetMutations)
	api.GET("/mutation/:id", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin,stocker"), MutationHandler.GetMutationByID)

	//stock adjustment
	api.POST("/stock-adjustments", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin,stocker"), stockAdjustmentHandler.Create)
	api.GET("/stock-adjustments", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin,stocker"), stockAdjustmentHandler.GetAll)
	api.GET("/stock-adjustments/:id", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin,stocker"), stockAdjustmentHandler.GetByID)
	// api.GET("/products/:product_id/stock-adjustments", middleware.AuthMiddleware(authService, userService), middleware.RequireRole("admin,stocker"), stockAdjustmentHandler.GetByProductID)

	r.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080

}
